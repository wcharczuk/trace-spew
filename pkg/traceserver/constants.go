package traceserver

// Headers
const (
	// HeaderTraceCount is a header containing the number of traces in the payload
	HeaderTraceCount  = "X-Datadog-Trace-Count"
	HeaderContainerID = "Datadog-Container-ID"
)

// ContentTypes
const (
	ContentTypeApplicationMessagePack = "application/msgpack"
)
