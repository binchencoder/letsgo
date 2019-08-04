package trace

import (
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const (
	// server需要的公共头信息.
	XUid            = "x-uid"         // 用户id
	XCid            = "x-cid"         // 企业id
	XAid            = "x-aid"         // 个人账户id
	XRequestId      = "x-request-id"  // 唯一请求id，跟踪日志使用
	XDid            = "x-did"         // 设备唯一id
	XAppVersion     = "x-app-version" // 客户端版本号， 如：3.4.0
	XSource         = "x-source"      // 请求来源。web,client
	XClient         = "x-client"      // 客户端资源号。如：mga、mip、win、mac
	XClientId       = "x-client-id"   // Open-plat required client_id
	XCorpCode       = "x-corp-code"   // Open-plat required 企业code
	S2SToken        = "x-s2s-token"   // Service to service token, used by grpc service to authenticate each other.
	Cookie          = "cookie"        // cookie
	JinSessionId    = "JINSESSIONID"  // JINSESSIONID,authn登陆成功后会set token到cookie的JINSESSIONID，authn后续接口读取JINSESSIONID获取token
	SetJinSessionId = "JINSESSIONID=" // JINSESSIONID with =

	// request source
	ResourceWeb          = "web"           // janus-gateway web端请求.
	ResourceClient       = "client"        // janus-gateway client端请求(包含windowns、mac、Android、ios4个客户端).
	ResourceOpenPlatform = "open-platform" // open-gateway requests.
	ResourceLdap         = "ldap"          // ldap-gateway requests.
)

// header info.
type Header struct {
	// 企业用户uid.
	Uid string
	// 企业用户cid.
	Cid string
	// 个人账号aid.
	Aid string
	// trace id is a string which uniquely identifies an organic request.
	TraceId string
	// 设备唯一id.
	Did string
	// 客户端版本号.
	AppVersion string
	// 请求来源.
	Source string
	// 客户端资源号.
	Client string
	// Open-plat required client_id.
	ClientId string
	// Open-plat required 企业code.
	CorpCode string
	// Service to service token.
	S2SToken string
}

// HeaderFromContext returns header info from incoming context.
func HeaderFromContext(incoming context.Context) *Header {
	md, ok := metadata.FromIncomingContext(incoming)
	if !ok {
		return nil
	}

	h := &Header{}
	if xs := md[XSource]; len(xs) > 0 {
		h.Source = xs[0]
	}
	if xr := md[XRequestId]; len(xr) > 0 {
		h.TraceId = xr[0]
	}
	if s2st := md[S2SToken]; len(s2st) > 0 {
		h.S2SToken = s2st[0]
	}

	if h.Source == ResourceOpenPlatform {
		h = openHeaderFromContext(h, md)
	} else if h.Source == ResourceLdap {
		h = ldapHeaderFromContext(h, md)
	} else {
		h = defaultHeaderFromContext(h, md)
	}

	return h
}

// defaultHeaderFromContext returns janus-gateway or grpc-grpc invoke header info from incoming context.
func defaultHeaderFromContext(h *Header, md metadata.MD) *Header {
	if xu := md[XUid]; len(xu) > 0 {
		h.Uid = xu[0]
	}
	if xc := md[XCid]; len(xc) > 0 {
		h.Cid = xc[0]
	}
	if xa := md[XAid]; len(xa) > 0 {
		h.Aid = xa[0]
	}
	if xcl := md[XClient]; len(xcl) > 0 {
		h.Client = xcl[0]
	}
	if xv := md[XAppVersion]; len(xv) > 0 {
		h.AppVersion = xv[0]
	}
	if xd := md[XDid]; len(xd) > 0 {
		h.Did = xd[0]
	}

	return h
}

// openHeaderFromContext returns open-gateway header info from incoming context.
func openHeaderFromContext(h *Header, md metadata.MD) *Header {
	if xcc := md[XCorpCode]; len(xcc) > 0 {
		h.CorpCode = xcc[0]
	}
	if xa := md[XClientId]; len(xa) > 0 {
		h.ClientId = xa[0]
	}
	if xc := md[XCid]; len(xc) > 0 {
		h.Cid = xc[0]
	}

	return h
}

// ldapHeaderFromContext returns ldap-gateway header info from incoming context.
func ldapHeaderFromContext(h *Header, md metadata.MD) *Header {
	// TODO(liww): 补充ldap-gateway的ctx信息.
	return h
}

// TokenFromContext returns token from incoming context.
func TokenFromContext(incoming context.Context) string {
	md, ok := metadata.FromIncomingContext(incoming)
	if !ok {
		return ""
	}

	if cookie := md[Cookie]; len(cookie) > 0 {
		arr := strings.Split(cookie[0], ";")
		for _, str := range arr {
			str = strings.TrimSpace(str)
			if strings.HasPrefix(str, JinSessionId) {
				return strings.TrimPrefix(str, SetJinSessionId)
			}
		}
	}

	return ""
}
