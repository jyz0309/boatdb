package mysql

const (
	ClientLongPassword uint32 = 1 << iota
	ClientFoundRows
	ClientLongFlag
	ClientConnectWithDB
	ClientNoSchema
	ClientCompress
	ClientODBC
	ClientLocalFiles
	ClientIgnoreSpace
	ClientProtocol41
	ClientInteractive
	ClientSSL
	ClientIgnoreSigpipe
	ClientTransactions
	ClientReserved
	ClientSecureConnection
	ClientMultiStatements
	ClientMultiResults
	ClientPSMultiResults
	ClientPluginAuth
	ClientConnectAtts
	ClientPluginAuthLenencClientData
)

var DefaultCapability = ClientLongPassword | ClientLongFlag |
	ClientConnectWithDB | ClientProtocol41 |
	ClientTransactions | ClientSecureConnection | ClientFoundRows |
	ClientMultiStatements | ClientMultiResults | ClientLocalFiles |
	ClientConnectAtts | ClientPluginAuth | ClientInteractive
