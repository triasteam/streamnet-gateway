package modules

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
)

type Path struct {
	Id						string		`json:"id"`
	DomainId				string		`json:"domain_id"`
	ReqMethod				string		`json:"req_method"`
	ReqPath					string		`json:"req_path"`
	SearchPath				string		`json:"search_path"`
	ReplacePath				string		`json:"replace_path"`
	CircuitBreakerRequest	int			`json:"circuit_breaker_request"`
	CircuitBreakerPercent	int			`json:"circuit_breaker_percent"`
	CircuitBreakerTimeout	int			`json:"circuit_breaker_timeout"`
	CircuitBreakerMsg		string		`json:"circuit_breaker_msg"`
	CircuitBreakerEnabled	bool		`json:"circuit_breaker_enabled"`
	CircuitBreakerForce		bool		`json:"circuit_breaker_force"`
	PrivateProxyEnabled		bool		`json:"private_proxy_enabled"`
	LbType					string		`json:"lb_type"`
	Targets					[]*Target	`json:"targets"`
	SetTime					string		`json:"set_time"`
}

func Paths(c *gin.Context) {
	domainId := c.Param("domain_id")
	if domainId == "" {
		mContext{c}.ErrorOP(DataParseError)
		return
	}

	rsp, err := pathsData(domainId)
	if err != nil {
		mContext{c}.ErrorOP(SystemError)
		return
	}
	var paths []*Path
	if rsp.Count > 0 {
		for _, kv := range rsp.Kvs {
			path := new(Path)
			err := json.Unmarshal(kv.Value, path)
			if err == nil {
				paths = append(paths, path)
			}
		}
		mContext{c}.SuccessOP(paths)
		return
	}
	mContext{c}.SuccessOP(make([]string, 0))
}

func PutPath(c *gin.Context) {
	domainId := c.Param("domain_id")
	reqMethod := c.PostForm("req_method")
	reqPath := c.PostForm("req_path")
	searchPath := c.PostForm("search_path")
	replacePath := c.PostForm("replace_path")
	cbRequest := c.PostForm("circuit_breaker_request")
	cbPercent := c.PostForm("circuit_breaker_percent")
	cbTimeout := c.PostForm("circuit_breaker_timeout")
	cbMsg := c.PostForm("circuit_breaker_msg")
	cbEnabled := c.PostForm("circuit_breaker_enabled")
	cbForce := c.PostForm("circuit_breaker_force")
	priProxyEnabled := c.PostForm("private_proxy_enabled")
	lbType := c.PostForm("lb_type")
	proxyTargets := c.PostForm("proxy_targets")

	if reqMethod == "" || reqPath == "" || domainId == "" {
		mContext{c}.ErrorOP(DataParseError)
		return
	}

	//检查代理目标数据
	var targets []*Target
	err := json.Unmarshal([]byte(proxyTargets), &targets)
	if err != nil {
		mContext{c}.ErrorOP(DataParseError)
		return
	}

	var pathId string
	pathId = c.Param("path_id")
	if pathId == "" {
		pathId = uuid.Must(uuid.NewV4()).String()
	}

	path := new(Path)
	path.Id = pathId
	path.DomainId = domainId
	path.ReqMethod = reqMethod
	path.ReqPath = reqPath
	path.SearchPath = searchPath
	path.ReplacePath = replacePath
	path.CircuitBreakerRequest,_ = strconv.Atoi(cbRequest)
	path.CircuitBreakerPercent,_ = strconv.Atoi(cbPercent)
	path.CircuitBreakerTimeout,_ = strconv.Atoi(cbTimeout)
	path.CircuitBreakerMsg = cbMsg
	path.CircuitBreakerEnabled,_ = strconv.ParseBool(cbEnabled)
	path.CircuitBreakerForce,_ = strconv.ParseBool(cbForce)
	path.LbType = lbType
	path.Targets = targets
	path.PrivateProxyEnabled,_ = strconv.ParseBool(priProxyEnabled)
	path.SetTime = time.Now().Format("2006/1/2 15:04:05")

	pathB, err := json.Marshal(path)
	if err != nil {
		mContext{c}.ErrorOP(DataParseError)
		return
	}

	err = putPath(domainId, path.Id, string(pathB))
	if err != nil {
		mContext{c}.ErrorOP(SystemError)
		return
	}
	mContext{c}.SuccessOP(path)
}

func GetPath(c *gin.Context) {
	domainId := c.Param("domain_id")
	pathId := c.Param("path_id")
	if domainId == "" || pathId == "" {
		mContext{c}.ErrorOP(DataParseError)
		return
	}
	rsp, err := pathData(domainId, pathId)
	if err != nil {
		mContext{c}.ErrorOP(DataParseError)
		return
	}

	if rsp.Count > 0 {
		path := new(Path)
		err := json.Unmarshal(rsp.Kvs[0].Value, path)
		if err != nil {
			mContext{c}.ErrorOP(DataParseError)
			return
		}
		mContext{c}.SuccessOP(path)
	} else {
		mContext{c}.SuccessOP(struct {}{})
	}
}

func DelPath(c *gin.Context) {
	domainId := c.Param("domain_id")
	pathId := c.Param("path_id")
	if domainId == "" || pathId == "" {
		mContext{c}.ErrorOP(DataParseError)
		return
	}

	deleted := delPath(domainId, pathId)
	if deleted  {
		mContext{c}.SuccessOP(nil)
		return
	}
	mContext{c}.ErrorOP(SystemError)
}