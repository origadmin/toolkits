package discovery

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/errors"
)

type ConsulConfig struct {
	// Address is the address of the Consul server
	Address string

	// Scheme is the URI scheme for the Consul server
	Scheme string

	// Prefix for URIs for when consul is behind an API gateway (reverse
	// proxy).  The API gateway must strip off the PathPrefix before
	// passing the request onto consul.
	PathPrefix string

	// Datacenter to use. If not provided, the default agent datacenter is used.
	Datacenter string

	// Transport is the Transport to use for the http client.
	Transport *http.Transport

	// HttpClient is the client to use. Default will be
	// used if not provided.
	HttpClient *http.Client

	// WaitTime limits how long a Watch will block. If not provided,
	// the agent default values will be used.
	WaitTime time.Duration

	// Token is used to provide a per-request ACL token
	// which overrides the agent's default token.
	Token string

	// TokenFile is a file containing the current token to use for this client.
	// If provided it is read once at startup and never again.
	TokenFile string

	// Namespace is the name of the namespace to send along for the request
	// when no other Namespace is present in the QueryOptions
	Namespace string

	// Partition is the name of the partition to send along for the request
	// when no other Partition is present in the QueryOptions
	Partition string

	TLSConfig TLSConfig
}

type TLSConfig struct {
	// Address is the optional address of the Consul server. The port, if any
	// will be removed from here and this will be set to the ServerName of the
	// resulting config.
	Address string

	// CAFile is the optional path to the CA certificate used for Consul
	// communication, defaults to the system bundle if not specified.
	CAFile string

	// CAPath is the optional path to a directory of CA certificates to use for
	// Consul communication, defaults to the system bundle if not specified.
	CAPath string

	// CAPem is the optional PEM-encoded CA certificate used for Consul
	// communication, defaults to the system bundle if not specified.
	CAPem []byte

	// CertFile is the optional path to the certificate for Consul
	// communication. If this is set then you need to also set KeyFile.
	CertFile string

	// CertPEM is the optional PEM-encoded certificate for Consul
	// communication. If this is set then you need to also set KeyPEM.
	CertPEM []byte

	// KeyFile is the optional path to the private key for Consul communication.
	// If this is set then you need to also set CertFile.
	KeyFile string

	// KeyPEM is the optional PEM-encoded private key for Consul communication.
	// If this is set then you need to also set CertPEM.
	KeyPEM []byte

	// InsecureSkipVerify if set to true will disable TLS host verification.
	InsecureSkipVerify bool
}

type ETCDConfig struct {
	// Endpoints is a list of URLs.
	Endpoints []string `json:"endpoints"`

	// AutoSyncInterval is the interval to update endpoints with its latest members.
	// 0 disables auto-sync. By default auto-sync is disabled.
	AutoSyncInterval time.Duration `json:"auto-sync-interval"`

	// DialTimeout is the timeout for failing to establish a connection.
	DialTimeout time.Duration `json:"dial-timeout"`

	// DialKeepAliveTime is the time after which client pings the server to see if
	// transport is alive.
	DialKeepAliveTime time.Duration `json:"dial-keep-alive-time"`

	// DialKeepAliveTimeout is the time that the client waits for a response for the
	// keep-alive probe. If the response is not received in this time, the connection is closed.
	DialKeepAliveTimeout time.Duration `json:"dial-keep-alive-timeout"`

	// MaxCallSendMsgSize is the client-side request send limit in bytes.
	// If 0, it defaults to 2.0 MiB (2 * 1024 * 1024).
	// Make sure that "MaxCallSendMsgSize" < server-side default send/recv limit.
	// ("--max-request-bytes" flag to etcd or "embed.Config.MaxRequestBytes").
	MaxCallSendMsgSize int

	// MaxCallRecvMsgSize is the client-side response receive limit.
	// If 0, it defaults to "math.MaxInt32", because range response can
	// easily exceed request send limits.
	// Make sure that "MaxCallRecvMsgSize" >= server-side default send/recv limit.
	// ("--max-request-bytes" flag to etcd or "embed.Config.MaxRequestBytes").
	MaxCallRecvMsgSize int

	// TLS holds the client secure credentials, if any.
	TLS *tls.Config

	// Username is a username for authentication.
	Username string `json:"username"`

	// Password is a password for authentication.
	Password string `json:"password"`

	// RejectOldCluster when set will refuse to create a client against an outdated cluster.
	RejectOldCluster bool `json:"reject-old-cluster"`

	// DialOptions is a list of dial options for the grpc client (e.g., for interceptors).
	// For example, pass "grpc.WithBlock()" to block until the underlying connection is up.
	// Without this, Dial returns immediately and connecting the server happens in background.
	DialOptions []grpc.DialOption

	// Context is the default client context; it can be used to cancel grpc dial out and
	// other operations that do not have an explicit context.
	Context context.Context

	// Logger sets client-side logger.
	// If nil, fallback to building LogConfig.
	Logger *zap.Logger

	// LogConfig configures client-side logger.
	// If nil, use the default logger.
	// TODO: configure gRPC logger
	LogConfig *zap.Config

	// PermitWithoutStream when set will allow client to send keepalive pings to server without any active streams(RPCs).
	PermitWithoutStream bool `json:"permit-without-stream"`

	// MaxUnaryRetries is the maximum number of retries for unary RPCs.
	MaxUnaryRetries uint `json:"max-unary-retries"`

	// BackoffWaitBetween is the wait time before retrying an RPC.
	BackoffWaitBetween time.Duration `json:"backoff-wait-between"`

	// BackoffJitterFraction is the jitter fraction to randomize backoff wait time.
	BackoffJitterFraction float64 `json:"backoff-jitter-fraction"`
}

type Config struct {
	Type   string
	Consul ConsulConfig
	ETCD   ETCDConfig
}

const (
	ErrNotFound = errors.String("not found")
)
