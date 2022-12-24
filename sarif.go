package sarif

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Address A physical or virtual address, or a range of addresses, in an 'addressable region' (memory or a binary file).
type Address struct {

	// The address expressed as a byte offset from the start of the addressable region.
	AbsoluteAddress int `json:"absoluteAddress,omitempty"`

	// A human-readable fully qualified name that is associated with the address.
	FullyQualifiedName string `json:"fullyQualifiedName,omitempty"`

	// The index within run.addresses of the cached object for this address.
	Index int `json:"index,omitempty"`

	// An open-ended string that identifies the address kind. 'data', 'function', 'header','instruction', 'module', 'page', 'section', 'segment', 'stack', 'stackFrame', 'table' are well-known values.
	Kind string `json:"kind,omitempty"`

	// The number of bytes in this range of addresses.
	Length int `json:"length,omitempty"`

	// A name that is associated with the address, e.g., '.text'.
	Name string `json:"name,omitempty"`

	// The byte offset of this address from the absolute or relative address of the parent object.
	OffsetFromParent int `json:"offsetFromParent,omitempty"`

	// The index within run.addresses of the parent object.
	ParentIndex int `json:"parentIndex,omitempty"`

	// Key/value pairs that provide additional information about the address.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The address expressed as a byte offset from the absolute address of the top-most parent object.
	RelativeAddress int `json:"relativeAddress,omitempty"`
}

// Artifact A single artifact. In some cases, this artifact might be nested within another artifact.
type Artifact struct {

	// The contents of the artifact.
	Contents *ArtifactContent `json:"contents,omitempty"`

	// A short description of the artifact.
	Description *Message `json:"description,omitempty"`

	// Specifies the encoding for an artifact object that refers to a text file.
	Encoding string `json:"encoding,omitempty"`

	// A dictionary, each of whose keys is the name of a hash function and each of whose values is the hashed value of the artifact produced by the specified hash function.
	Hashes map[string]string `json:"hashes,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the artifact was most recently modified. See "Date/time properties" in the SARIF spec for the required format.
	LastModifiedTimeUtc string `json:"lastModifiedTimeUtc,omitempty"`

	// The length of the artifact in bytes.
	Length int `json:"length,omitempty"`

	// The location of the artifact.
	Location *ArtifactLocation `json:"location,omitempty"`

	// The MIME type (RFC 2045) of the artifact.
	MimeType string `json:"mimeType,omitempty"`

	// The offset in bytes of the artifact within its containing artifact.
	Offset int `json:"offset,omitempty"`

	// Identifies the index of the immediate parent of the artifact, if this artifact is nested.
	ParentIndex int `json:"parentIndex,omitempty"`

	// Key/value pairs that provide additional information about the artifact.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The role or roles played by the artifact in the analysis.
	Roles []string `json:"roles,omitempty"`

	// Specifies the source language for any artifact object that refers to a text file that contains source code.
	SourceLanguage string `json:"sourceLanguage,omitempty"`
}

// ArtifactChange A change to a single artifact.
type ArtifactChange struct {

	// The location of the artifact to change.
	ArtifactLocation *ArtifactLocation `json:"artifactLocation"`

	// Key/value pairs that provide additional information about the change.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of replacement objects, each of which represents the replacement of a single region in a single artifact specified by 'artifactLocation'.
	Replacements []*Replacement `json:"replacements"`
}

// ArtifactContent Represents the contents of an artifact.
type ArtifactContent struct {

	// MIME Base64-encoded content from a binary artifact, or from a text artifact in its original encoding.
	Binary string `json:"binary,omitempty"`

	// Key/value pairs that provide additional information about the artifact content.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An alternate rendered representation of the artifact (e.g., a decompiled representation of a binary region).
	Rendered *MultiformatMessageString `json:"rendered,omitempty"`

	// UTF-8-encoded content from a text artifact.
	Text string `json:"text,omitempty"`
}

// ArtifactLocation Specifies the location of an artifact.
type ArtifactLocation struct {

	// A short description of the artifact location.
	Description *Message `json:"description,omitempty"`

	// The index within the run artifacts array of the artifact object associated with the artifact location.
	Index int `json:"index,omitempty"`

	// Key/value pairs that provide additional information about the artifact location.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A string containing a valid relative or absolute URI.
	Uri string `json:"uri,omitempty"`

	// A string which indirectly specifies the absolute URI with respect to which a relative URI in the "uri" property is interpreted.
	UriBaseId string `json:"uriBaseId,omitempty"`
}

// Attachment An artifact relevant to a result.
type Attachment struct {

	// The location of the attachment.
	ArtifactLocation *ArtifactLocation `json:"artifactLocation"`

	// A message describing the role played by the attachment.
	Description *Message `json:"description,omitempty"`

	// Key/value pairs that provide additional information about the attachment.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of rectangles specifying areas of interest within the image.
	Rectangles []*Rectangle `json:"rectangles,omitempty"`

	// An array of regions of interest within the attachment.
	Regions []*Region `json:"regions,omitempty"`
}

// CodeFlow A set of threadFlows which together describe a pattern of code execution relevant to detecting a result.
type CodeFlow struct {

	// A message relevant to the code flow.
	Message *Message `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the code flow.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of one or more unique threadFlow objects, each of which describes the progress of a program through a thread of execution.
	ThreadFlows []*ThreadFlow `json:"threadFlows"`
}

// ConfigurationOverride Information about how a specific rule or notification was reconfigured at runtime.
type ConfigurationOverride struct {

	// Specifies how the rule or notification was configured during the scan.
	Configuration *ReportingConfiguration `json:"configuration"`

	// A reference used to locate the descriptor whose configuration was overridden.
	Descriptor *ReportingDescriptorReference `json:"descriptor"`

	// Key/value pairs that provide additional information about the configuration override.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// Conversion Describes how a converter transformed the output of a static analysis tool from the analysis tool's native output format into the SARIF format.
type Conversion struct {

	// The locations of the analysis tool's per-run log files.
	AnalysisToolLogFiles []*ArtifactLocation `json:"analysisToolLogFiles,omitempty"`

	// An invocation object that describes the invocation of the converter.
	Invocation *Invocation `json:"invocation,omitempty"`

	// Key/value pairs that provide additional information about the conversion.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A tool object that describes the converter.
	Tool *Tool `json:"tool"`
}

// Edge Represents a directed edge in a graph.
type Edge struct {

	// A string that uniquely identifies the edge within its graph.
	Id string `json:"id"`

	// A short description of the edge.
	Label *Message `json:"label,omitempty"`

	// Key/value pairs that provide additional information about the edge.
	Properties *PropertyBag `json:"properties,omitempty"`

	// Identifies the source node (the node at which the edge starts).
	SourceNodeId string `json:"sourceNodeId"`

	// Identifies the target node (the node at which the edge ends).
	TargetNodeId string `json:"targetNodeId"`
}

// EdgeTraversal Represents the traversal of a single edge during a graph traversal.
type EdgeTraversal struct {

	// Identifies the edge being traversed.
	EdgeId string `json:"edgeId"`

	// The values of relevant expressions after the edge has been traversed.
	FinalState map[string]*MultiformatMessageString `json:"finalState,omitempty"`

	// A message to display to the user as the edge is traversed.
	Message *Message `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the edge traversal.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The number of edge traversals necessary to return from a nested graph.
	StepOverEdgeCount int `json:"stepOverEdgeCount,omitempty"`
}

// Exception Describes a runtime exception encountered during the execution of an analysis tool.
type Exception struct {

	// An array of exception objects each of which is considered a cause of this exception.
	InnerExceptions []*Exception `json:"innerExceptions,omitempty"`

	// A string that identifies the kind of exception, for example, the fully qualified type name of an object that was thrown, or the symbolic name of a signal.
	Kind string `json:"kind,omitempty"`

	// A message that describes the exception.
	Message string `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the exception.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The sequence of function calls leading to the exception.
	Stack *Stack `json:"stack,omitempty"`
}

// ExternalProperties The top-level element of an external property file.
type ExternalProperties struct {

	// Addresses that will be merged with a separate run.
	Addresses []*Address `json:"addresses,omitempty"`

	// An array of artifact objects that will be merged with a separate run.
	Artifacts []*Artifact `json:"artifacts,omitempty"`

	// A conversion object that will be merged with a separate run.
	Conversion *Conversion `json:"conversion,omitempty"`

	// The analysis tool object that will be merged with a separate run.
	Driver *ToolComponent `json:"driver,omitempty"`

	// Tool extensions that will be merged with a separate run.
	Extensions []*ToolComponent `json:"extensions,omitempty"`

	// Key/value pairs that provide additional information that will be merged with a separate run.
	ExternalizedProperties *PropertyBag `json:"externalizedProperties,omitempty"`

	// An array of graph objects that will be merged with a separate run.
	Graphs []*Graph `json:"graphs,omitempty"`

	// A stable, unique identifer for this external properties object, in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// Describes the invocation of the analysis tool that will be merged with a separate run.
	Invocations []*Invocation `json:"invocations,omitempty"`

	// An array of logical locations such as namespaces, types or functions that will be merged with a separate run.
	LogicalLocations []*LogicalLocation `json:"logicalLocations,omitempty"`

	// Tool policies that will be merged with a separate run.
	Policies []*ToolComponent `json:"policies,omitempty"`

	// Key/value pairs that provide additional information about the external properties.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of result objects that will be merged with a separate run.
	Results []*Result `json:"results,omitempty"`

	// A stable, unique identifer for the run associated with this external properties object, in the form of a GUID.
	RunGuid string `json:"runGuid,omitempty"`

	// The URI of the JSON schema corresponding to the version of the external property file format.
	Schema string `json:"schema,omitempty"`

	// Tool taxonomies that will be merged with a separate run.
	Taxonomies []*ToolComponent `json:"taxonomies,omitempty"`

	// An array of threadFlowLocation objects that will be merged with a separate run.
	ThreadFlowLocations []*ThreadFlowLocation `json:"threadFlowLocations,omitempty"`

	// Tool translations that will be merged with a separate run.
	Translations []*ToolComponent `json:"translations,omitempty"`

	// The SARIF format version of this external properties object.
	Version string `json:"version,omitempty"`

	// Requests that will be merged with a separate run.
	WebRequests []*WebRequest `json:"webRequests,omitempty"`

	// Responses that will be merged with a separate run.
	WebResponses []*WebResponse `json:"webResponses,omitempty"`
}

// ExternalPropertyFileReference Contains information that enables a SARIF consumer to locate the external property file that contains the value of an externalized property associated with the run.
type ExternalPropertyFileReference struct {

	// A stable, unique identifer for the external property file in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// A non-negative integer specifying the number of items contained in the external property file.
	ItemCount int `json:"itemCount,omitempty"`

	// The location of the external property file.
	Location *ArtifactLocation `json:"location,omitempty"`

	// Key/value pairs that provide additional information about the external property file.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// ExternalPropertyFileReferences References to external property files that should be inlined with the content of a root log file.
type ExternalPropertyFileReferences struct {

	// An array of external property files containing run.addresses arrays to be merged with the root log file.
	Addresses []*ExternalPropertyFileReference `json:"addresses,omitempty"`

	// An array of external property files containing run.artifacts arrays to be merged with the root log file.
	Artifacts []*ExternalPropertyFileReference `json:"artifacts,omitempty"`

	// An external property file containing a run.conversion object to be merged with the root log file.
	Conversion *ExternalPropertyFileReference `json:"conversion,omitempty"`

	// An external property file containing a run.driver object to be merged with the root log file.
	Driver *ExternalPropertyFileReference `json:"driver,omitempty"`

	// An array of external property files containing run.extensions arrays to be merged with the root log file.
	Extensions []*ExternalPropertyFileReference `json:"extensions,omitempty"`

	// An external property file containing a run.properties object to be merged with the root log file.
	ExternalizedProperties *ExternalPropertyFileReference `json:"externalizedProperties,omitempty"`

	// An array of external property files containing a run.graphs object to be merged with the root log file.
	Graphs []*ExternalPropertyFileReference `json:"graphs,omitempty"`

	// An array of external property files containing run.invocations arrays to be merged with the root log file.
	Invocations []*ExternalPropertyFileReference `json:"invocations,omitempty"`

	// An array of external property files containing run.logicalLocations arrays to be merged with the root log file.
	LogicalLocations []*ExternalPropertyFileReference `json:"logicalLocations,omitempty"`

	// An array of external property files containing run.policies arrays to be merged with the root log file.
	Policies []*ExternalPropertyFileReference `json:"policies,omitempty"`

	// Key/value pairs that provide additional information about the external property files.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of external property files containing run.results arrays to be merged with the root log file.
	Results []*ExternalPropertyFileReference `json:"results,omitempty"`

	// An array of external property files containing run.taxonomies arrays to be merged with the root log file.
	Taxonomies []*ExternalPropertyFileReference `json:"taxonomies,omitempty"`

	// An array of external property files containing run.threadFlowLocations arrays to be merged with the root log file.
	ThreadFlowLocations []*ExternalPropertyFileReference `json:"threadFlowLocations,omitempty"`

	// An array of external property files containing run.translations arrays to be merged with the root log file.
	Translations []*ExternalPropertyFileReference `json:"translations,omitempty"`

	// An array of external property files containing run.requests arrays to be merged with the root log file.
	WebRequests []*ExternalPropertyFileReference `json:"webRequests,omitempty"`

	// An array of external property files containing run.responses arrays to be merged with the root log file.
	WebResponses []*ExternalPropertyFileReference `json:"webResponses,omitempty"`
}

// Fix A proposed fix for the problem represented by a result object. A fix specifies a set of artifacts to modify. For each artifact, it specifies a set of bytes to remove, and provides a set of new bytes to replace them.
type Fix struct {

	// One or more artifact changes that comprise a fix for a result.
	ArtifactChanges []*ArtifactChange `json:"artifactChanges"`

	// A message that describes the proposed fix, enabling viewers to present the proposed change to an end user.
	Description *Message `json:"description,omitempty"`

	// Key/value pairs that provide additional information about the fix.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// Graph A network of nodes and directed edges that describes some aspect of the structure of the code (for example, a call graph).
type Graph struct {

	// A description of the graph.
	Description *Message `json:"description,omitempty"`

	// An array of edge objects representing the edges of the graph.
	Edges []*Edge `json:"edges,omitempty"`

	// An array of node objects representing the nodes of the graph.
	Nodes []*Node `json:"nodes,omitempty"`

	// Key/value pairs that provide additional information about the graph.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// GraphTraversal Represents a path through a graph.
type GraphTraversal struct {

	// A description of this graph traversal.
	Description *Message `json:"description,omitempty"`

	// The sequences of edges traversed by this graph traversal.
	EdgeTraversals []*EdgeTraversal `json:"edgeTraversals,omitempty"`

	// Values of relevant expressions at the start of the graph traversal that remain constant for the graph traversal.
	ImmutableState map[string]*MultiformatMessageString `json:"immutableState,omitempty"`

	// Values of relevant expressions at the start of the graph traversal that may change during graph traversal.
	InitialState map[string]*MultiformatMessageString `json:"initialState,omitempty"`

	// Key/value pairs that provide additional information about the graph traversal.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The index within the result.graphs to be associated with the result.
	ResultGraphIndex int `json:"resultGraphIndex,omitempty"`

	// The index within the run.graphs to be associated with the result.
	RunGraphIndex int `json:"runGraphIndex,omitempty"`
}

// Invocation The runtime environment of the analysis tool run.
type Invocation struct {

	// The account that ran the analysis tool.
	Account string `json:"account,omitempty"`

	// An array of strings, containing in order the command line arguments passed to the tool from the operating system.
	Arguments []string `json:"arguments,omitempty"`

	// The command line used to invoke the tool.
	CommandLine string `json:"commandLine,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the run ended. See "Date/time properties" in the SARIF spec for the required format.
	EndTimeUtc string `json:"endTimeUtc,omitempty"`

	// The environment variables associated with the analysis tool process, expressed as key/value pairs.
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	// An absolute URI specifying the location of the analysis tool's executable.
	ExecutableLocation *ArtifactLocation `json:"executableLocation,omitempty"`

	// Specifies whether the tool's execution completed successfully.
	ExecutionSuccessful bool `json:"executionSuccessful"`

	// The process exit code.
	ExitCode int `json:"exitCode,omitempty"`

	// The reason for the process exit.
	ExitCodeDescription string `json:"exitCodeDescription,omitempty"`

	// The name of the signal that caused the process to exit.
	ExitSignalName string `json:"exitSignalName,omitempty"`

	// The numeric value of the signal that caused the process to exit.
	ExitSignalNumber int `json:"exitSignalNumber,omitempty"`

	// The machine that hosted the analysis tool run.
	Machine string `json:"machine,omitempty"`

	// An array of configurationOverride objects that describe notifications related runtime overrides.
	NotificationConfigurationOverrides []*ConfigurationOverride `json:"notificationConfigurationOverrides,omitempty"`

	// The process id for the analysis tool run.
	ProcessId int `json:"processId,omitempty"`

	// The reason given by the operating system that the process failed to start.
	ProcessStartFailureMessage string `json:"processStartFailureMessage,omitempty"`

	// Key/value pairs that provide additional information about the invocation.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The locations of any response files specified on the tool's command line.
	ResponseFiles []*ArtifactLocation `json:"responseFiles,omitempty"`

	// An array of configurationOverride objects that describe rules related runtime overrides.
	RuleConfigurationOverrides []*ConfigurationOverride `json:"ruleConfigurationOverrides,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the run started. See "Date/time properties" in the SARIF spec for the required format.
	StartTimeUtc string `json:"startTimeUtc,omitempty"`

	// A file containing the standard error stream from the process that was invoked.
	Stderr *ArtifactLocation `json:"stderr,omitempty"`

	// A file containing the standard input stream to the process that was invoked.
	Stdin *ArtifactLocation `json:"stdin,omitempty"`

	// A file containing the standard output stream from the process that was invoked.
	Stdout *ArtifactLocation `json:"stdout,omitempty"`

	// A file containing the interleaved standard output and standard error stream from the process that was invoked.
	StdoutStderr *ArtifactLocation `json:"stdoutStderr,omitempty"`

	// A list of conditions detected by the tool that are relevant to the tool's configuration.
	ToolConfigurationNotifications []*Notification `json:"toolConfigurationNotifications,omitempty"`

	// A list of runtime conditions detected by the tool during the analysis.
	ToolExecutionNotifications []*Notification `json:"toolExecutionNotifications,omitempty"`

	// The working directory for the analysis tool run.
	WorkingDirectory *ArtifactLocation `json:"workingDirectory,omitempty"`
}

// Location A location within a programming artifact.
type Location struct {

	// A set of regions relevant to the location.
	Annotations []*Region `json:"annotations,omitempty"`

	// Value that distinguishes this location from all other locations within a single result object.
	Id int `json:"id,omitempty"`

	// The logical locations associated with the result.
	LogicalLocations []*LogicalLocation `json:"logicalLocations,omitempty"`

	// A message relevant to the location.
	Message *Message `json:"message,omitempty"`

	// Identifies the artifact and region.
	PhysicalLocation *PhysicalLocation `json:"physicalLocation,omitempty"`

	// Key/value pairs that provide additional information about the location.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of objects that describe relationships between this location and others.
	Relationships []*LocationRelationship `json:"relationships,omitempty"`
}

// LocationRelationship Information about the relation of one location to another.
type LocationRelationship struct {

	// A description of the location relationship.
	Description *Message `json:"description,omitempty"`

	// A set of distinct strings that categorize the relationship. Well-known kinds include 'includes', 'isIncludedBy' and 'relevant'.
	Kinds []string `json:"kinds,omitempty"`

	// Key/value pairs that provide additional information about the location relationship.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A reference to the related location.
	Target int `json:"target"`
}

// LogicalLocation A logical location of a construct that produced a result.
type LogicalLocation struct {

	// The machine-readable name for the logical location, such as a mangled function name provided by a C++ compiler that encodes calling convention, return type and other details along with the function name.
	DecoratedName string `json:"decoratedName,omitempty"`

	// The human-readable fully qualified name of the logical location.
	FullyQualifiedName string `json:"fullyQualifiedName,omitempty"`

	// The index within the logical locations array.
	Index int `json:"index,omitempty"`

	// The type of construct this logical location component refers to. Should be one of 'function', 'member', 'module', 'namespace', 'parameter', 'resource', 'returnType', 'type', 'variable', 'object', 'array', 'property', 'value', 'element', 'text', 'attribute', 'comment', 'declaration', 'dtd' or 'processingInstruction', if any of those accurately describe the construct.
	Kind string `json:"kind,omitempty"`

	// Identifies the construct in which the result occurred. For example, this property might contain the name of a class or a method.
	Name string `json:"name,omitempty"`

	// Identifies the index of the immediate parent of the construct in which the result was detected. For example, this property might point to a logical location that represents the namespace that holds a type.
	ParentIndex int `json:"parentIndex,omitempty"`

	// Key/value pairs that provide additional information about the logical location.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// Message Encapsulates a message intended to be read by the end user.
type Message struct {

	// An array of strings to substitute into the message string.
	Arguments []string `json:"arguments,omitempty"`

	// The identifier for this message.
	Id string `json:"id,omitempty"`

	// A Markdown message string.
	Markdown string `json:"markdown,omitempty"`

	// Key/value pairs that provide additional information about the message.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A plain text message string.
	Text string `json:"text,omitempty"`
}

// MultiformatMessageString A message string or message format string rendered in multiple formats.
type MultiformatMessageString struct {

	// A Markdown message string or format string.
	Markdown string `json:"markdown,omitempty"`

	// Key/value pairs that provide additional information about the message.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A plain text message string or format string.
	Text string `json:"text"`
}

// Node Represents a node in a graph.
type Node struct {

	// Array of child nodes.
	Children []*Node `json:"children,omitempty"`

	// A string that uniquely identifies the node within its graph.
	Id string `json:"id"`

	// A short description of the node.
	Label *Message `json:"label,omitempty"`

	// A code location associated with the node.
	Location *Location `json:"location,omitempty"`

	// Key/value pairs that provide additional information about the node.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// Notification Describes a condition relevant to the tool itself, as opposed to being relevant to a target being analyzed by the tool.
type Notification struct {

	// A reference used to locate the rule descriptor associated with this notification.
	AssociatedRule *ReportingDescriptorReference `json:"associatedRule,omitempty"`

	// A reference used to locate the descriptor relevant to this notification.
	Descriptor *ReportingDescriptorReference `json:"descriptor,omitempty"`

	// The runtime exception, if any, relevant to this notification.
	Exception *Exception `json:"exception,omitempty"`

	// A value specifying the severity level of the notification.
	Level string `json:"level,omitempty"`

	// The locations relevant to this notification.
	Locations []*Location `json:"locations,omitempty"`

	// A message that describes the condition that was encountered.
	Message *Message `json:"message"`

	// Key/value pairs that provide additional information about the notification.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The thread identifier of the code that generated the notification.
	ThreadId int `json:"threadId,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the analysis tool generated the notification.
	TimeUtc string `json:"timeUtc,omitempty"`
}

// PhysicalLocation A physical location relevant to a result. Specifies a reference to a programming artifact together with a range of bytes or characters within that artifact.
type PhysicalLocation struct {

	// The address of the location.
	Address *Address `json:"address,omitempty"`

	// The location of the artifact.
	ArtifactLocation *ArtifactLocation `json:"artifactLocation,omitempty"`

	// Specifies a portion of the artifact that encloses the region. Allows a viewer to display additional context around the region.
	ContextRegion *Region `json:"contextRegion,omitempty"`

	// Key/value pairs that provide additional information about the physical location.
	Properties *PropertyBag `json:"properties,omitempty"`

	// Specifies a portion of the artifact.
	Region *Region `json:"region,omitempty"`
}

// PropertyBag Key/value pairs that provide additional information about the object.
type PropertyBag struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`

	// A set of distinct strings that provide additional information.
	Tags []string `json:"tags,omitempty"`
}

// Rectangle An area within an image.
type Rectangle struct {

	// The Y coordinate of the bottom edge of the rectangle, measured in the image's natural units.
	Bottom float64 `json:"bottom,omitempty"`

	// The X coordinate of the left edge of the rectangle, measured in the image's natural units.
	Left float64 `json:"left,omitempty"`

	// A message relevant to the rectangle.
	Message *Message `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the rectangle.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The X coordinate of the right edge of the rectangle, measured in the image's natural units.
	Right float64 `json:"right,omitempty"`

	// The Y coordinate of the top edge of the rectangle, measured in the image's natural units.
	Top float64 `json:"top,omitempty"`
}

// Region A region within an artifact where a result was detected.
type Region struct {

	// The length of the region in bytes.
	ByteLength int `json:"byteLength,omitempty"`

	// The zero-based offset from the beginning of the artifact of the first byte in the region.
	ByteOffset int `json:"byteOffset,omitempty"`

	// The length of the region in characters.
	CharLength int `json:"charLength,omitempty"`

	// The zero-based offset from the beginning of the artifact of the first character in the region.
	CharOffset int `json:"charOffset,omitempty"`

	// The column number of the character following the end of the region.
	EndColumn int `json:"endColumn,omitempty"`

	// The line number of the last character in the region.
	EndLine int `json:"endLine,omitempty"`

	// A message relevant to the region.
	Message *Message `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the region.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The portion of the artifact contents within the specified region.
	Snippet *ArtifactContent `json:"snippet,omitempty"`

	// Specifies the source language, if any, of the portion of the artifact specified by the region object.
	SourceLanguage string `json:"sourceLanguage,omitempty"`

	// The column number of the first character in the region.
	StartColumn int `json:"startColumn,omitempty"`

	// The line number of the first character in the region.
	StartLine int `json:"startLine,omitempty"`
}

// Replacement The replacement of a single region of an artifact.
type Replacement struct {

	// The region of the artifact to delete.
	DeletedRegion *Region `json:"deletedRegion"`

	// The content to insert at the location specified by the 'deletedRegion' property.
	InsertedContent *ArtifactContent `json:"insertedContent,omitempty"`

	// Key/value pairs that provide additional information about the replacement.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// ReportingConfiguration Information about a rule or notification that can be configured at runtime.
type ReportingConfiguration struct {

	// Specifies whether the report may be produced during the scan.
	Enabled bool `json:"enabled,omitempty"`

	// Specifies the failure level for the report.
	Level string `json:"level,omitempty"`

	// Contains configuration information specific to a report.
	Parameters *PropertyBag `json:"parameters,omitempty"`

	// Key/value pairs that provide additional information about the reporting configuration.
	Properties *PropertyBag `json:"properties,omitempty"`

	// Specifies the relative priority of the report. Used for analysis output only.
	Rank float64 `json:"rank,omitempty"`
}

// ReportingDescriptor Metadata that describes a specific report produced by the tool, as part of the analysis it provides or its runtime reporting.
type ReportingDescriptor struct {

	// Default reporting configuration information.
	DefaultConfiguration *ReportingConfiguration `json:"defaultConfiguration,omitempty"`

	// An array of unique identifies in the form of a GUID by which this report was known in some previous version of the analysis tool.
	DeprecatedGuids []string `json:"deprecatedGuids,omitempty"`

	// An array of stable, opaque identifiers by which this report was known in some previous version of the analysis tool.
	DeprecatedIds []string `json:"deprecatedIds,omitempty"`

	// An array of readable identifiers by which this report was known in some previous version of the analysis tool.
	DeprecatedNames []string `json:"deprecatedNames,omitempty"`

	// A description of the report. Should, as far as possible, provide details sufficient to enable resolution of any problem indicated by the result.
	FullDescription *MultiformatMessageString `json:"fullDescription,omitempty"`

	// A unique identifer for the reporting descriptor in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// Provides the primary documentation for the report, useful when there is no online documentation.
	Help *MultiformatMessageString `json:"help,omitempty"`

	// A URI where the primary documentation for the report can be found.
	HelpUri string `json:"helpUri,omitempty"`

	// A stable, opaque identifier for the report.
	Id string `json:"id"`

	// A set of name/value pairs with arbitrary names. Each value is a multiformatMessageString object, which holds message strings in plain text and (optionally) Markdown format. The strings can include placeholders, which can be used to construct a message in combination with an arbitrary number of additional string arguments.
	MessageStrings map[string]*MultiformatMessageString `json:"messageStrings,omitempty"`

	// A report identifier that is understandable to an end user.
	Name string `json:"name,omitempty"`

	// Key/value pairs that provide additional information about the report.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of objects that describe relationships between this reporting descriptor and others.
	Relationships []*ReportingDescriptorRelationship `json:"relationships,omitempty"`

	// A concise description of the report. Should be a single sentence that is understandable when visible space is limited to a single line of text.
	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`
}

// ReportingDescriptorReference Information about how to locate a relevant reporting descriptor.
type ReportingDescriptorReference struct {

	// A guid that uniquely identifies the descriptor.
	Guid string `json:"guid,omitempty"`

	// The id of the descriptor.
	Id string `json:"id,omitempty"`

	// The index into an array of descriptors in toolComponent.ruleDescriptors, toolComponent.notificationDescriptors, or toolComponent.taxonomyDescriptors, depending on context.
	Index int `json:"index,omitempty"`

	// Key/value pairs that provide additional information about the reporting descriptor reference.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A reference used to locate the toolComponent associated with the descriptor.
	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`
}

// ReportingDescriptorRelationship Information about the relation of one reporting descriptor to another.
type ReportingDescriptorRelationship struct {

	// A description of the reporting descriptor relationship.
	Description *Message `json:"description,omitempty"`

	// A set of distinct strings that categorize the relationship. Well-known kinds include 'canPrecede', 'canFollow', 'willPrecede', 'willFollow', 'superset', 'subset', 'equal', 'disjoint', 'relevant', and 'incomparable'.
	Kinds []string `json:"kinds,omitempty"`

	// Key/value pairs that provide additional information about the reporting descriptor reference.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A reference to the related reporting descriptor.
	Target *ReportingDescriptorReference `json:"target"`
}

// Result A result produced by an analysis tool.
type Result struct {

	// Identifies the artifact that the analysis tool was instructed to scan. This need not be the same as the artifact where the result actually occurred.
	AnalysisTarget *ArtifactLocation `json:"analysisTarget,omitempty"`

	// A set of artifacts relevant to the result.
	Attachments []*Attachment `json:"attachments,omitempty"`

	// The state of a result relative to a baseline of a previous run.
	BaselineState string `json:"baselineState,omitempty"`

	// An array of 'codeFlow' objects relevant to the result.
	CodeFlows []*CodeFlow `json:"codeFlows,omitempty"`

	// A stable, unique identifier for the equivalence class of logically identical results to which this result belongs, in the form of a GUID.
	CorrelationGuid string `json:"correlationGuid,omitempty"`

	// A set of strings each of which individually defines a stable, unique identity for the result.
	Fingerprints map[string]string `json:"fingerprints,omitempty"`

	// An array of 'fix' objects, each of which represents a proposed fix to the problem indicated by the result.
	Fixes []*Fix `json:"fixes,omitempty"`

	// An array of one or more unique 'graphTraversal' objects.
	GraphTraversals []*GraphTraversal `json:"graphTraversals,omitempty"`

	// An array of zero or more unique graph objects associated with the result.
	Graphs []*Graph `json:"graphs,omitempty"`

	// A stable, unique identifer for the result in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// An absolute URI at which the result can be viewed.
	HostedViewerUri string `json:"hostedViewerUri,omitempty"`

	// A value that categorizes results by evaluation state.
	Kind string `json:"kind,omitempty"`

	// A value specifying the severity level of the result.
	Level string `json:"level,omitempty"`

	// The set of locations where the result was detected. Specify only one location unless the problem indicated by the result can only be corrected by making a change at every specified location.
	Locations []*Location `json:"locations,omitempty"`

	// A message that describes the result. The first sentence of the message only will be displayed when visible space is limited.
	Message *Message `json:"message"`

	// A positive integer specifying the number of times this logically unique result was observed in this run.
	OccurrenceCount int `json:"occurrenceCount,omitempty"`

	// A set of strings that contribute to the stable, unique identity of the result.
	PartialFingerprints map[string]string `json:"partialFingerprints,omitempty"`

	// Key/value pairs that provide additional information about the result.
	Properties *PropertyBag `json:"properties,omitempty"`

	// Information about how and when the result was detected.
	Provenance *ResultProvenance `json:"provenance,omitempty"`

	// A number representing the priority or importance of the result.
	Rank float64 `json:"rank,omitempty"`

	// A set of locations relevant to this result.
	RelatedLocations []*Location `json:"relatedLocations,omitempty"`

	// A reference used to locate the rule descriptor relevant to this result.
	Rule *ReportingDescriptorReference `json:"rule,omitempty"`

	// The stable, unique identifier of the rule, if any, to which this notification is relevant. This member can be used to retrieve rule metadata from the rules dictionary, if it exists.
	RuleId string `json:"ruleId,omitempty"`

	// The index within the tool component rules array of the rule object associated with this result.
	RuleIndex int `json:"ruleIndex,omitempty"`

	// An array of 'stack' objects relevant to the result.
	Stacks []*Stack `json:"stacks,omitempty"`

	// A set of suppressions relevant to this result.
	Suppressions []*Suppression `json:"suppressions,omitempty"`

	// An array of references to taxonomy reporting descriptors that are applicable to the result.
	Taxa []*ReportingDescriptorReference `json:"taxa,omitempty"`

	// A web request associated with this result.
	WebRequest *WebRequest `json:"webRequest,omitempty"`

	// A web response associated with this result.
	WebResponse *WebResponse `json:"webResponse,omitempty"`

	// The URIs of the work items associated with this result.
	WorkItemUris []string `json:"workItemUris,omitempty"`
}

// ResultProvenance Contains information about how and when a result was detected.
type ResultProvenance struct {

	// An array of physicalLocation objects which specify the portions of an analysis tool's output that a converter transformed into the result.
	ConversionSources []*PhysicalLocation `json:"conversionSources,omitempty"`

	// A GUID-valued string equal to the automationDetails.guid property of the run in which the result was first detected.
	FirstDetectionRunGuid string `json:"firstDetectionRunGuid,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the result was first detected. See "Date/time properties" in the SARIF spec for the required format.
	FirstDetectionTimeUtc string `json:"firstDetectionTimeUtc,omitempty"`

	// The index within the run.invocations array of the invocation object which describes the tool invocation that detected the result.
	InvocationIndex int `json:"invocationIndex,omitempty"`

	// A GUID-valued string equal to the automationDetails.guid property of the run in which the result was most recently detected.
	LastDetectionRunGuid string `json:"lastDetectionRunGuid,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which the result was most recently detected. See "Date/time properties" in the SARIF spec for the required format.
	LastDetectionTimeUtc string `json:"lastDetectionTimeUtc,omitempty"`

	// Key/value pairs that provide additional information about the result.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// Run Describes a single run of an analysis tool, and contains the reported output of that run.
type Run struct {

	// Addresses associated with this run instance, if any.
	Addresses []*Address `json:"addresses,omitempty"`

	// An array of artifact objects relevant to the run.
	Artifacts []*Artifact `json:"artifacts,omitempty"`

	// Automation details that describe this run.
	AutomationDetails *RunAutomationDetails `json:"automationDetails,omitempty"`

	// The 'guid' property of a previous SARIF 'run' that comprises the baseline that was used to compute result 'baselineState' properties for the run.
	BaselineGuid string `json:"baselineGuid,omitempty"`

	// Specifies the unit in which the tool measures columns.
	ColumnKind string `json:"columnKind,omitempty"`

	// A conversion object that describes how a converter transformed an analysis tool's native reporting format into the SARIF format.
	Conversion *Conversion `json:"conversion,omitempty"`

	// Specifies the default encoding for any artifact object that refers to a text file.
	DefaultEncoding string `json:"defaultEncoding,omitempty"`

	// Specifies the default source language for any artifact object that refers to a text file that contains source code.
	DefaultSourceLanguage string `json:"defaultSourceLanguage,omitempty"`

	// References to external property files that should be inlined with the content of a root log file.
	ExternalPropertyFileReferences *ExternalPropertyFileReferences `json:"externalPropertyFileReferences,omitempty"`

	// An array of zero or more unique graph objects associated with the run.
	Graphs []*Graph `json:"graphs,omitempty"`

	// Describes the invocation of the analysis tool.
	Invocations []*Invocation `json:"invocations,omitempty"`

	// The language of the messages emitted into the log file during this run (expressed as an ISO 639-1 two-letter lowercase culture code) and an optional region (expressed as an ISO 3166-1 two-letter uppercase subculture code associated with a country or region). The casing is recommended but not required (in order for this data to conform to RFC5646).
	Language string `json:"language,omitempty"`

	// An array of logical locations such as namespaces, types or functions.
	LogicalLocations []*LogicalLocation `json:"logicalLocations,omitempty"`

	// An ordered list of character sequences that were treated as line breaks when computing region information for the run.
	NewlineSequences []string `json:"newlineSequences,omitempty"`

	// The artifact location specified by each uriBaseId symbol on the machine where the tool originally ran.
	OriginalUriBaseIds map[string]*ArtifactLocation `json:"originalUriBaseIds,omitempty"`

	// Contains configurations that may potentially override both reportingDescriptor.defaultConfiguration (the tool's default severities) and invocation.configurationOverrides (severities established at run-time from the command line).
	Policies []*ToolComponent `json:"policies,omitempty"`

	// Key/value pairs that provide additional information about the run.
	Properties *PropertyBag `json:"properties,omitempty"`

	// An array of strings used to replace sensitive information in a redaction-aware property.
	RedactionTokens []string `json:"redactionTokens,omitempty"`

	// The set of results contained in an SARIF log. The results array can be omitted when a run is solely exporting rules metadata. It must be present (but may be empty) if a log file represents an actual scan.
	Results []*Result `json:"results,omitempty"`

	// Automation details that describe the aggregate of runs to which this run belongs.
	RunAggregates []*RunAutomationDetails `json:"runAggregates,omitempty"`

	// A specialLocations object that defines locations of special significance to SARIF consumers.
	SpecialLocations *SpecialLocations `json:"specialLocations,omitempty"`

	// An array of toolComponent objects relevant to a taxonomy in which results are categorized.
	Taxonomies []*ToolComponent `json:"taxonomies,omitempty"`

	// An array of threadFlowLocation objects cached at run level.
	ThreadFlowLocations []*ThreadFlowLocation `json:"threadFlowLocations,omitempty"`

	// Information about the tool or tool pipeline that generated the results in this run. A run can only contain results produced by a single tool or tool pipeline. A run can aggregate results from multiple log files, as long as context around the tool run (tool command-line arguments and the like) is identical for all aggregated files.
	Tool *Tool `json:"tool"`

	// The set of available translations of the localized data provided by the tool.
	Translations []*ToolComponent `json:"translations,omitempty"`

	// Specifies the revision in version control of the artifacts that were scanned.
	VersionControlProvenance []*VersionControlDetails `json:"versionControlProvenance,omitempty"`

	// An array of request objects cached at run level.
	WebRequests []*WebRequest `json:"webRequests,omitempty"`

	// An array of response objects cached at run level.
	WebResponses []*WebResponse `json:"webResponses,omitempty"`
}

// RunAutomationDetails Information that describes a run's identity and role within an engineering system process.
type RunAutomationDetails struct {

	// A stable, unique identifier for the equivalence class of runs to which this object's containing run object belongs in the form of a GUID.
	CorrelationGuid string `json:"correlationGuid,omitempty"`

	// A description of the identity and role played within the engineering system by this object's containing run object.
	Description *Message `json:"description,omitempty"`

	// A stable, unique identifer for this object's containing run object in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// A hierarchical string that uniquely identifies this object's containing run object.
	Id string `json:"id,omitempty"`

	// Key/value pairs that provide additional information about the run automation details.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// SpecialLocations Defines locations of special significance to SARIF consumers.
type SpecialLocations struct {

	// Provides a suggestion to SARIF consumers to display file paths relative to the specified location.
	DisplayBase *ArtifactLocation `json:"displayBase,omitempty"`

	// Key/value pairs that provide additional information about the special locations.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// Stack A call stack that is relevant to a result.
type Stack struct {

	// An array of stack frames that represents a sequence of calls, rendered in reverse chronological order, that comprise the call stack.
	Frames []*StackFrame `json:"frames"`

	// A message relevant to this call stack.
	Message *Message `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the stack.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// StackFrame A function call within a stack trace.
type StackFrame struct {

	// The location to which this stack frame refers.
	Location *Location `json:"location,omitempty"`

	// The name of the module that contains the code of this stack frame.
	Module string `json:"module,omitempty"`

	// The parameters of the call that is executing.
	Parameters []string `json:"parameters,omitempty"`

	// Key/value pairs that provide additional information about the stack frame.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The thread identifier of the stack frame.
	ThreadId int `json:"threadId,omitempty"`
}

// SARIF Static Analysis Results Format (SARIF) Version 2.1.0 JSON Schema: a standard format for the output of static analysis tools.
type SARIF struct {

	// References to external property files that share data between runs.
	InlineExternalProperties []*ExternalProperties `json:"inlineExternalProperties,omitempty"`

	// Key/value pairs that provide additional information about the log file.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The set of runs contained in this log file.
	Runs []*Run `json:"runs"`

	// The URI of the JSON schema corresponding to the version.
	Schema string `json:"$schema,omitempty"`

	// The SARIF format version of this log file.
	Version string `json:"version"`
}

// Suppression A suppression that is relevant to a result.
type Suppression struct {

	// A stable, unique identifer for the suprression in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// A string representing the justification for the suppression.
	Justification string `json:"justification,omitempty"`

	// A string that indicates where the suppression is persisted.
	Kind string `json:"kind"`

	// Identifies the location associated with the suppression.
	Location *Location `json:"location,omitempty"`

	// Key/value pairs that provide additional information about the suppression.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A string that indicates the state of the suppression.
	State string `json:"state,omitempty"`
}

// ThreadFlow Describes a sequence of code locations that specify a path through a single thread of execution such as an operating system or fiber.
type ThreadFlow struct {

	// An string that uniquely identifies the threadFlow within the codeFlow in which it occurs.
	Id string `json:"id,omitempty"`

	// Values of relevant expressions at the start of the thread flow that remain constant.
	ImmutableState map[string]*MultiformatMessageString `json:"immutableState,omitempty"`

	// Values of relevant expressions at the start of the thread flow that may change during thread flow execution.
	InitialState map[string]*MultiformatMessageString `json:"initialState,omitempty"`

	// A temporally ordered array of 'threadFlowLocation' objects, each of which describes a location visited by the tool while producing the result.
	Locations []*ThreadFlowLocation `json:"locations"`

	// A message relevant to the thread flow.
	Message *Message `json:"message,omitempty"`

	// Key/value pairs that provide additional information about the thread flow.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// ThreadFlowLocation A location visited by an analysis tool while simulating or monitoring the execution of a program.
type ThreadFlowLocation struct {

	// An integer representing the temporal order in which execution reached this location.
	ExecutionOrder int `json:"executionOrder,omitempty"`

	// The Coordinated Universal Time (UTC) date and time at which this location was executed.
	ExecutionTimeUtc string `json:"executionTimeUtc,omitempty"`

	// Specifies the importance of this location in understanding the code flow in which it occurs. The order from most to least important is "essential", "important", "unimportant". Default: "important".
	Importance string `json:"importance,omitempty"`

	// The index within the run threadFlowLocations array.
	Index int `json:"index,omitempty"`

	// A set of distinct strings that categorize the thread flow location. Well-known kinds include 'acquire', 'release', 'enter', 'exit', 'call', 'return', 'branch', 'implicit', 'false', 'true', 'caution', 'danger', 'unknown', 'unreachable', 'taint', 'function', 'handler', 'lock', 'memory', 'resource', 'scope' and 'value'.
	Kinds []string `json:"kinds,omitempty"`

	// The code location.
	Location *Location `json:"location,omitempty"`

	// The name of the module that contains the code that is executing.
	Module string `json:"module,omitempty"`

	// An integer representing a containment hierarchy within the thread flow.
	NestingLevel int `json:"nestingLevel,omitempty"`

	// Key/value pairs that provide additional information about the threadflow location.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The call stack leading to this location.
	Stack *Stack `json:"stack,omitempty"`

	// A dictionary, each of whose keys specifies a variable or expression, the associated value of which represents the variable or expression value. For an annotation of kind 'continuation', for example, this dictionary might hold the current assumed values of a set of global variables.
	State map[string]*MultiformatMessageString `json:"state,omitempty"`

	// An array of references to rule or taxonomy reporting descriptors that are applicable to the thread flow location.
	Taxa []*ReportingDescriptorReference `json:"taxa,omitempty"`

	// A web request associated with this thread flow location.
	WebRequest *WebRequest `json:"webRequest,omitempty"`

	// A web response associated with this thread flow location.
	WebResponse *WebResponse `json:"webResponse,omitempty"`
}

// Tool The analysis tool that was run.
type Tool struct {

	// The analysis tool that was run.
	Driver *ToolComponent `json:"driver"`

	// Tool extensions that contributed to or reconfigured the analysis tool that was run.
	Extensions []*ToolComponent `json:"extensions,omitempty"`

	// Key/value pairs that provide additional information about the tool.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// ToolComponent A component, such as a plug-in or the driver, of the analysis tool that was run.
type ToolComponent struct {

	// The component which is strongly associated with this component. For a translation, this refers to the component which has been translated. For an extension, this is the driver that provides the extension's plugin model.
	AssociatedComponent *ToolComponentReference `json:"associatedComponent,omitempty"`

	// The kinds of data contained in this object.
	Contents string `json:"contents,omitempty"`

	// The binary version of the tool component's primary executable file expressed as four non-negative integers separated by a period (for operating systems that express file versions in this way).
	DottedQuadFileVersion string `json:"dottedQuadFileVersion,omitempty"`

	// The absolute URI from which the tool component can be downloaded.
	DownloadUri string `json:"downloadUri,omitempty"`

	// A comprehensive description of the tool component.
	FullDescription *MultiformatMessageString `json:"fullDescription,omitempty"`

	// The name of the tool component along with its version and any other useful identifying information, such as its locale.
	FullName string `json:"fullName,omitempty"`

	// A dictionary, each of whose keys is a resource identifier and each of whose values is a multiformatMessageString object, which holds message strings in plain text and (optionally) Markdown format. The strings can include placeholders, which can be used to construct a message in combination with an arbitrary number of additional string arguments.
	GlobalMessageStrings map[string]*MultiformatMessageString `json:"globalMessageStrings,omitempty"`

	// A unique identifer for the tool component in the form of a GUID.
	Guid string `json:"guid,omitempty"`

	// The absolute URI at which information about this version of the tool component can be found.
	InformationUri string `json:"informationUri,omitempty"`

	// Specifies whether this object contains a complete definition of the localizable and/or non-localizable data for this component, as opposed to including only data that is relevant to the results persisted to this log file.
	IsComprehensive bool `json:"isComprehensive,omitempty"`

	// The language of the messages emitted into the log file during this run (expressed as an ISO 639-1 two-letter lowercase language code) and an optional region (expressed as an ISO 3166-1 two-letter uppercase subculture code associated with a country or region). The casing is recommended but not required (in order for this data to conform to RFC5646).
	Language string `json:"language,omitempty"`

	// The semantic version of the localized strings defined in this component; maintained by components that provide translations.
	LocalizedDataSemanticVersion string `json:"localizedDataSemanticVersion,omitempty"`

	// An array of the artifactLocation objects associated with the tool component.
	Locations []*ArtifactLocation `json:"locations,omitempty"`

	// The minimum value of localizedDataSemanticVersion required in translations consumed by this component; used by components that consume translations.
	MinimumRequiredLocalizedDataSemanticVersion string `json:"minimumRequiredLocalizedDataSemanticVersion,omitempty"`

	// The name of the tool component.
	Name string `json:"name"`

	// An array of reportingDescriptor objects relevant to the notifications related to the configuration and runtime execution of the tool component.
	Notifications []*ReportingDescriptor `json:"notifications,omitempty"`

	// The organization or company that produced the tool component.
	Organization string `json:"organization,omitempty"`

	// A product suite to which the tool component belongs.
	Product string `json:"product,omitempty"`

	// A localizable string containing the name of the suite of products to which the tool component belongs.
	ProductSuite string `json:"productSuite,omitempty"`

	// Key/value pairs that provide additional information about the tool component.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A string specifying the UTC date (and optionally, the time) of the component's release.
	ReleaseDateUtc string `json:"releaseDateUtc,omitempty"`

	// An array of reportingDescriptor objects relevant to the analysis performed by the tool component.
	Rules []*ReportingDescriptor `json:"rules,omitempty"`

	// The tool component version in the format specified by Semantic Versioning 2.0.
	SemanticVersion string `json:"semanticVersion,omitempty"`

	// A brief description of the tool component.
	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`

	// An array of toolComponentReference objects to declare the taxonomies supported by the tool component.
	SupportedTaxonomies []*ToolComponentReference `json:"supportedTaxonomies,omitempty"`

	// An array of reportingDescriptor objects relevant to the definitions of both standalone and tool-defined taxonomies.
	Taxa []*ReportingDescriptor `json:"taxa,omitempty"`

	// Translation metadata, required for a translation, not populated by other component types.
	TranslationMetadata *TranslationMetadata `json:"translationMetadata,omitempty"`

	// The tool component version, in whatever format the component natively provides.
	Version string `json:"version,omitempty"`
}

// ToolComponentReference Identifies a particular toolComponent object, either the driver or an extension.
type ToolComponentReference struct {

	// The 'guid' property of the referenced toolComponent.
	Guid string `json:"guid,omitempty"`

	// An index into the referenced toolComponent in tool.extensions.
	Index int `json:"index,omitempty"`

	// The 'name' property of the referenced toolComponent.
	Name string `json:"name,omitempty"`

	// Key/value pairs that provide additional information about the toolComponentReference.
	Properties *PropertyBag `json:"properties,omitempty"`
}

// TranslationMetadata Provides additional metadata related to translation.
type TranslationMetadata struct {

	// The absolute URI from which the translation metadata can be downloaded.
	DownloadUri string `json:"downloadUri,omitempty"`

	// A comprehensive description of the translation metadata.
	FullDescription *MultiformatMessageString `json:"fullDescription,omitempty"`

	// The full name associated with the translation metadata.
	FullName string `json:"fullName,omitempty"`

	// The absolute URI from which information related to the translation metadata can be downloaded.
	InformationUri string `json:"informationUri,omitempty"`

	// The name associated with the translation metadata.
	Name string `json:"name"`

	// Key/value pairs that provide additional information about the translation metadata.
	Properties *PropertyBag `json:"properties,omitempty"`

	// A brief description of the translation metadata.
	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`
}

// VersionControlDetails Specifies the information necessary to retrieve a desired revision from a version control system.
type VersionControlDetails struct {

	// A Coordinated Universal Time (UTC) date and time that can be used to synchronize an enlistment to the state of the repository at that time.
	AsOfTimeUtc string `json:"asOfTimeUtc,omitempty"`

	// The name of a branch containing the revision.
	Branch string `json:"branch,omitempty"`

	// The location in the local file system to which the root of the repository was mapped at the time of the analysis.
	MappedTo *ArtifactLocation `json:"mappedTo,omitempty"`

	// Key/value pairs that provide additional information about the version control details.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The absolute URI of the repository.
	RepositoryUri string `json:"repositoryUri"`

	// A string that uniquely and permanently identifies the revision within the repository.
	RevisionId string `json:"revisionId,omitempty"`

	// A tag that has been applied to the revision.
	RevisionTag string `json:"revisionTag,omitempty"`
}

// WebRequest Describes an HTTP request.
type WebRequest struct {

	// The body of the request.
	Body *ArtifactContent `json:"body,omitempty"`

	// The request headers.
	Headers map[string]string `json:"headers,omitempty"`

	// The index within the run.webRequests array of the request object associated with this result.
	Index int `json:"index,omitempty"`

	// The HTTP method. Well-known values are 'GET', 'PUT', 'POST', 'DELETE', 'PATCH', 'HEAD', 'OPTIONS', 'TRACE', 'CONNECT'.
	Method string `json:"method,omitempty"`

	// The request parameters.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Key/value pairs that provide additional information about the request.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The request protocol. Example: 'http'.
	Protocol string `json:"protocol,omitempty"`

	// The target of the request.
	Target string `json:"target,omitempty"`

	// The request version. Example: '1.1'.
	Version string `json:"version,omitempty"`
}

// WebResponse Describes the response to an HTTP request.
type WebResponse struct {

	// The body of the response.
	Body *ArtifactContent `json:"body,omitempty"`

	// The response headers.
	Headers map[string]string `json:"headers,omitempty"`

	// The index within the run.webResponses array of the response object associated with this result.
	Index int `json:"index,omitempty"`

	// Specifies whether a response was received from the server.
	NoResponseReceived bool `json:"noResponseReceived,omitempty"`

	// Key/value pairs that provide additional information about the response.
	Properties *PropertyBag `json:"properties,omitempty"`

	// The response protocol. Example: 'http'.
	Protocol string `json:"protocol,omitempty"`

	// The response reason. Example: 'Not found'.
	ReasonPhrase string `json:"reasonPhrase,omitempty"`

	// The response status code. Example: 451.
	StatusCode int `json:"statusCode,omitempty"`

	// The response version. Example: '1.1'.
	Version string `json:"version,omitempty"`
}

func (strct *Address) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "absoluteAddress" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"absoluteAddress\": ")
	if tmp, err := json.Marshal(strct.AbsoluteAddress); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullyQualifiedName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullyQualifiedName\": ")
	if tmp, err := json.Marshal(strct.FullyQualifiedName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "length" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"length\": ")
	if tmp, err := json.Marshal(strct.Length); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "offsetFromParent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"offsetFromParent\": ")
	if tmp, err := json.Marshal(strct.OffsetFromParent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parentIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parentIndex\": ")
	if tmp, err := json.Marshal(strct.ParentIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "relativeAddress" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"relativeAddress\": ")
	if tmp, err := json.Marshal(strct.RelativeAddress); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Address) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "absoluteAddress":
			if err := json.Unmarshal([]byte(v), &strct.AbsoluteAddress); err != nil {
				return err
			}
		case "fullyQualifiedName":
			if err := json.Unmarshal([]byte(v), &strct.FullyQualifiedName); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "kind":
			if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
				return err
			}
		case "length":
			if err := json.Unmarshal([]byte(v), &strct.Length); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		case "offsetFromParent":
			if err := json.Unmarshal([]byte(v), &strct.OffsetFromParent); err != nil {
				return err
			}
		case "parentIndex":
			if err := json.Unmarshal([]byte(v), &strct.ParentIndex); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "relativeAddress":
			if err := json.Unmarshal([]byte(v), &strct.RelativeAddress); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Artifact) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "contents" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"contents\": ")
	if tmp, err := json.Marshal(strct.Contents); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "encoding" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"encoding\": ")
	if tmp, err := json.Marshal(strct.Encoding); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "hashes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"hashes\": ")
	if tmp, err := json.Marshal(strct.Hashes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "lastModifiedTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"lastModifiedTimeUtc\": ")
	if tmp, err := json.Marshal(strct.LastModifiedTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "length" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"length\": ")
	if tmp, err := json.Marshal(strct.Length); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "mimeType" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"mimeType\": ")
	if tmp, err := json.Marshal(strct.MimeType); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "offset" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"offset\": ")
	if tmp, err := json.Marshal(strct.Offset); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parentIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parentIndex\": ")
	if tmp, err := json.Marshal(strct.ParentIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "roles" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"roles\": ")
	if tmp, err := json.Marshal(strct.Roles); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "sourceLanguage" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sourceLanguage\": ")
	if tmp, err := json.Marshal(strct.SourceLanguage); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Artifact) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "contents":
			if err := json.Unmarshal([]byte(v), &strct.Contents); err != nil {
				return err
			}
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "encoding":
			if err := json.Unmarshal([]byte(v), &strct.Encoding); err != nil {
				return err
			}
		case "hashes":
			if err := json.Unmarshal([]byte(v), &strct.Hashes); err != nil {
				return err
			}
		case "lastModifiedTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.LastModifiedTimeUtc); err != nil {
				return err
			}
		case "length":
			if err := json.Unmarshal([]byte(v), &strct.Length); err != nil {
				return err
			}
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
		case "mimeType":
			if err := json.Unmarshal([]byte(v), &strct.MimeType); err != nil {
				return err
			}
		case "offset":
			if err := json.Unmarshal([]byte(v), &strct.Offset); err != nil {
				return err
			}
		case "parentIndex":
			if err := json.Unmarshal([]byte(v), &strct.ParentIndex); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "roles":
			if err := json.Unmarshal([]byte(v), &strct.Roles); err != nil {
				return err
			}
		case "sourceLanguage":
			if err := json.Unmarshal([]byte(v), &strct.SourceLanguage); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ArtifactChange) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ArtifactLocation" field is required
	if strct.ArtifactLocation == nil {
		return nil, errors.New("artifactLocation is a required field")
	}
	// Marshal the "artifactLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifactLocation\": ")
	if tmp, err := json.Marshal(strct.ArtifactLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Replacements" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "replacements" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"replacements\": ")
	if tmp, err := json.Marshal(strct.Replacements); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ArtifactChange) UnmarshalJSON(b []byte) error {
	artifactLocationReceived := false
	replacementsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "artifactLocation":
			if err := json.Unmarshal([]byte(v), &strct.ArtifactLocation); err != nil {
				return err
			}
			artifactLocationReceived = true
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "replacements":
			if err := json.Unmarshal([]byte(v), &strct.Replacements); err != nil {
				return err
			}
			replacementsReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if artifactLocation (a required property) was received
	if !artifactLocationReceived {
		return errors.New("\"artifactLocation\" is required but was not present")
	}
	// check if replacements (a required property) was received
	if !replacementsReceived {
		return errors.New("\"replacements\" is required but was not present")
	}
	return nil
}

func (strct *ArtifactContent) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "binary" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"binary\": ")
	if tmp, err := json.Marshal(strct.Binary); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rendered" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rendered\": ")
	if tmp, err := json.Marshal(strct.Rendered); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "text" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"text\": ")
	if tmp, err := json.Marshal(strct.Text); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ArtifactContent) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "binary":
			if err := json.Unmarshal([]byte(v), &strct.Binary); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "rendered":
			if err := json.Unmarshal([]byte(v), &strct.Rendered); err != nil {
				return err
			}
		case "text":
			if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ArtifactLocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "uri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"uri\": ")
	if tmp, err := json.Marshal(strct.Uri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "uriBaseId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"uriBaseId\": ")
	if tmp, err := json.Marshal(strct.UriBaseId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ArtifactLocation) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "uri":
			if err := json.Unmarshal([]byte(v), &strct.Uri); err != nil {
				return err
			}
		case "uriBaseId":
			if err := json.Unmarshal([]byte(v), &strct.UriBaseId); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Attachment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ArtifactLocation" field is required
	if strct.ArtifactLocation == nil {
		return nil, errors.New("artifactLocation is a required field")
	}
	// Marshal the "artifactLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifactLocation\": ")
	if tmp, err := json.Marshal(strct.ArtifactLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rectangles" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rectangles\": ")
	if tmp, err := json.Marshal(strct.Rectangles); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "regions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"regions\": ")
	if tmp, err := json.Marshal(strct.Regions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Attachment) UnmarshalJSON(b []byte) error {
	artifactLocationReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "artifactLocation":
			if err := json.Unmarshal([]byte(v), &strct.ArtifactLocation); err != nil {
				return err
			}
			artifactLocationReceived = true
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "rectangles":
			if err := json.Unmarshal([]byte(v), &strct.Rectangles); err != nil {
				return err
			}
		case "regions":
			if err := json.Unmarshal([]byte(v), &strct.Regions); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if artifactLocation (a required property) was received
	if !artifactLocationReceived {
		return errors.New("\"artifactLocation\" is required but was not present")
	}
	return nil
}

func (strct *CodeFlow) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ThreadFlows" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "threadFlows" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"threadFlows\": ")
	if tmp, err := json.Marshal(strct.ThreadFlows); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *CodeFlow) UnmarshalJSON(b []byte) error {
	threadFlowsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "threadFlows":
			if err := json.Unmarshal([]byte(v), &strct.ThreadFlows); err != nil {
				return err
			}
			threadFlowsReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if threadFlows (a required property) was received
	if !threadFlowsReceived {
		return errors.New("\"threadFlows\" is required but was not present")
	}
	return nil
}

func (strct *ConfigurationOverride) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Configuration" field is required
	if strct.Configuration == nil {
		return nil, errors.New("configuration is a required field")
	}
	// Marshal the "configuration" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configuration\": ")
	if tmp, err := json.Marshal(strct.Configuration); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Descriptor" field is required
	if strct.Descriptor == nil {
		return nil, errors.New("descriptor is a required field")
	}
	// Marshal the "descriptor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"descriptor\": ")
	if tmp, err := json.Marshal(strct.Descriptor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ConfigurationOverride) UnmarshalJSON(b []byte) error {
	configurationReceived := false
	descriptorReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "configuration":
			if err := json.Unmarshal([]byte(v), &strct.Configuration); err != nil {
				return err
			}
			configurationReceived = true
		case "descriptor":
			if err := json.Unmarshal([]byte(v), &strct.Descriptor); err != nil {
				return err
			}
			descriptorReceived = true
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if configuration (a required property) was received
	if !configurationReceived {
		return errors.New("\"configuration\" is required but was not present")
	}
	// check if descriptor (a required property) was received
	if !descriptorReceived {
		return errors.New("\"descriptor\" is required but was not present")
	}
	return nil
}

func (strct *Conversion) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "analysisToolLogFiles" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"analysisToolLogFiles\": ")
	if tmp, err := json.Marshal(strct.AnalysisToolLogFiles); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "invocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"invocation\": ")
	if tmp, err := json.Marshal(strct.Invocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Tool" field is required
	if strct.Tool == nil {
		return nil, errors.New("tool is a required field")
	}
	// Marshal the "tool" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tool\": ")
	if tmp, err := json.Marshal(strct.Tool); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Conversion) UnmarshalJSON(b []byte) error {
	toolReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "analysisToolLogFiles":
			if err := json.Unmarshal([]byte(v), &strct.AnalysisToolLogFiles); err != nil {
				return err
			}
		case "invocation":
			if err := json.Unmarshal([]byte(v), &strct.Invocation); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "tool":
			if err := json.Unmarshal([]byte(v), &strct.Tool); err != nil {
				return err
			}
			toolReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if tool (a required property) was received
	if !toolReceived {
		return errors.New("\"tool\" is required but was not present")
	}
	return nil
}

func (strct *Edge) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Id" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "label" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"label\": ")
	if tmp, err := json.Marshal(strct.Label); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "SourceNodeId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "sourceNodeId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sourceNodeId\": ")
	if tmp, err := json.Marshal(strct.SourceNodeId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "TargetNodeId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "targetNodeId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"targetNodeId\": ")
	if tmp, err := json.Marshal(strct.TargetNodeId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Edge) UnmarshalJSON(b []byte) error {
	idReceived := false
	sourceNodeIdReceived := false
	targetNodeIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
			idReceived = true
		case "label":
			if err := json.Unmarshal([]byte(v), &strct.Label); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "sourceNodeId":
			if err := json.Unmarshal([]byte(v), &strct.SourceNodeId); err != nil {
				return err
			}
			sourceNodeIdReceived = true
		case "targetNodeId":
			if err := json.Unmarshal([]byte(v), &strct.TargetNodeId); err != nil {
				return err
			}
			targetNodeIdReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if id (a required property) was received
	if !idReceived {
		return errors.New("\"id\" is required but was not present")
	}
	// check if sourceNodeId (a required property) was received
	if !sourceNodeIdReceived {
		return errors.New("\"sourceNodeId\" is required but was not present")
	}
	// check if targetNodeId (a required property) was received
	if !targetNodeIdReceived {
		return errors.New("\"targetNodeId\" is required but was not present")
	}
	return nil
}

func (strct *EdgeTraversal) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "EdgeId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "edgeId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"edgeId\": ")
	if tmp, err := json.Marshal(strct.EdgeId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "finalState" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"finalState\": ")
	if tmp, err := json.Marshal(strct.FinalState); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stepOverEdgeCount" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stepOverEdgeCount\": ")
	if tmp, err := json.Marshal(strct.StepOverEdgeCount); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EdgeTraversal) UnmarshalJSON(b []byte) error {
	edgeIdReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "edgeId":
			if err := json.Unmarshal([]byte(v), &strct.EdgeId); err != nil {
				return err
			}
			edgeIdReceived = true
		case "finalState":
			if err := json.Unmarshal([]byte(v), &strct.FinalState); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "stepOverEdgeCount":
			if err := json.Unmarshal([]byte(v), &strct.StepOverEdgeCount); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if edgeId (a required property) was received
	if !edgeIdReceived {
		return errors.New("\"edgeId\" is required but was not present")
	}
	return nil
}

func (strct *Exception) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "innerExceptions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"innerExceptions\": ")
	if tmp, err := json.Marshal(strct.InnerExceptions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stack" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stack\": ")
	if tmp, err := json.Marshal(strct.Stack); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Exception) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "innerExceptions":
			if err := json.Unmarshal([]byte(v), &strct.InnerExceptions); err != nil {
				return err
			}
		case "kind":
			if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "stack":
			if err := json.Unmarshal([]byte(v), &strct.Stack); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ExternalProperties) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "addresses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"addresses\": ")
	if tmp, err := json.Marshal(strct.Addresses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "artifacts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifacts\": ")
	if tmp, err := json.Marshal(strct.Artifacts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "conversion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"conversion\": ")
	if tmp, err := json.Marshal(strct.Conversion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "extensions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extensions\": ")
	if tmp, err := json.Marshal(strct.Extensions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "externalizedProperties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"externalizedProperties\": ")
	if tmp, err := json.Marshal(strct.ExternalizedProperties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "graphs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"graphs\": ")
	if tmp, err := json.Marshal(strct.Graphs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "invocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"invocations\": ")
	if tmp, err := json.Marshal(strct.Invocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "logicalLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"logicalLocations\": ")
	if tmp, err := json.Marshal(strct.LogicalLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "policies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"policies\": ")
	if tmp, err := json.Marshal(strct.Policies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "results" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"results\": ")
	if tmp, err := json.Marshal(strct.Results); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "runGuid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"runGuid\": ")
	if tmp, err := json.Marshal(strct.RunGuid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "schema" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"schema\": ")
	if tmp, err := json.Marshal(strct.Schema); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "taxonomies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"taxonomies\": ")
	if tmp, err := json.Marshal(strct.Taxonomies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "threadFlowLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"threadFlowLocations\": ")
	if tmp, err := json.Marshal(strct.ThreadFlowLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "translations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"translations\": ")
	if tmp, err := json.Marshal(strct.Translations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webRequests" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webRequests\": ")
	if tmp, err := json.Marshal(strct.WebRequests); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webResponses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webResponses\": ")
	if tmp, err := json.Marshal(strct.WebResponses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ExternalProperties) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "addresses":
			if err := json.Unmarshal([]byte(v), &strct.Addresses); err != nil {
				return err
			}
		case "artifacts":
			if err := json.Unmarshal([]byte(v), &strct.Artifacts); err != nil {
				return err
			}
		case "conversion":
			if err := json.Unmarshal([]byte(v), &strct.Conversion); err != nil {
				return err
			}
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
		case "extensions":
			if err := json.Unmarshal([]byte(v), &strct.Extensions); err != nil {
				return err
			}
		case "externalizedProperties":
			if err := json.Unmarshal([]byte(v), &strct.ExternalizedProperties); err != nil {
				return err
			}
		case "graphs":
			if err := json.Unmarshal([]byte(v), &strct.Graphs); err != nil {
				return err
			}
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "invocations":
			if err := json.Unmarshal([]byte(v), &strct.Invocations); err != nil {
				return err
			}
		case "logicalLocations":
			if err := json.Unmarshal([]byte(v), &strct.LogicalLocations); err != nil {
				return err
			}
		case "policies":
			if err := json.Unmarshal([]byte(v), &strct.Policies); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "results":
			if err := json.Unmarshal([]byte(v), &strct.Results); err != nil {
				return err
			}
		case "runGuid":
			if err := json.Unmarshal([]byte(v), &strct.RunGuid); err != nil {
				return err
			}
		case "schema":
			if err := json.Unmarshal([]byte(v), &strct.Schema); err != nil {
				return err
			}
		case "taxonomies":
			if err := json.Unmarshal([]byte(v), &strct.Taxonomies); err != nil {
				return err
			}
		case "threadFlowLocations":
			if err := json.Unmarshal([]byte(v), &strct.ThreadFlowLocations); err != nil {
				return err
			}
		case "translations":
			if err := json.Unmarshal([]byte(v), &strct.Translations); err != nil {
				return err
			}
		case "version":
			if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
				return err
			}
		case "webRequests":
			if err := json.Unmarshal([]byte(v), &strct.WebRequests); err != nil {
				return err
			}
		case "webResponses":
			if err := json.Unmarshal([]byte(v), &strct.WebResponses); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ExternalPropertyFileReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "itemCount" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"itemCount\": ")
	if tmp, err := json.Marshal(strct.ItemCount); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ExternalPropertyFileReference) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "itemCount":
			if err := json.Unmarshal([]byte(v), &strct.ItemCount); err != nil {
				return err
			}
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ExternalPropertyFileReferences) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "addresses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"addresses\": ")
	if tmp, err := json.Marshal(strct.Addresses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "artifacts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifacts\": ")
	if tmp, err := json.Marshal(strct.Artifacts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "conversion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"conversion\": ")
	if tmp, err := json.Marshal(strct.Conversion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "extensions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extensions\": ")
	if tmp, err := json.Marshal(strct.Extensions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "externalizedProperties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"externalizedProperties\": ")
	if tmp, err := json.Marshal(strct.ExternalizedProperties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "graphs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"graphs\": ")
	if tmp, err := json.Marshal(strct.Graphs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "invocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"invocations\": ")
	if tmp, err := json.Marshal(strct.Invocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "logicalLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"logicalLocations\": ")
	if tmp, err := json.Marshal(strct.LogicalLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "policies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"policies\": ")
	if tmp, err := json.Marshal(strct.Policies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "results" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"results\": ")
	if tmp, err := json.Marshal(strct.Results); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "taxonomies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"taxonomies\": ")
	if tmp, err := json.Marshal(strct.Taxonomies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "threadFlowLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"threadFlowLocations\": ")
	if tmp, err := json.Marshal(strct.ThreadFlowLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "translations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"translations\": ")
	if tmp, err := json.Marshal(strct.Translations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webRequests" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webRequests\": ")
	if tmp, err := json.Marshal(strct.WebRequests); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webResponses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webResponses\": ")
	if tmp, err := json.Marshal(strct.WebResponses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ExternalPropertyFileReferences) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "addresses":
			if err := json.Unmarshal([]byte(v), &strct.Addresses); err != nil {
				return err
			}
		case "artifacts":
			if err := json.Unmarshal([]byte(v), &strct.Artifacts); err != nil {
				return err
			}
		case "conversion":
			if err := json.Unmarshal([]byte(v), &strct.Conversion); err != nil {
				return err
			}
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
		case "extensions":
			if err := json.Unmarshal([]byte(v), &strct.Extensions); err != nil {
				return err
			}
		case "externalizedProperties":
			if err := json.Unmarshal([]byte(v), &strct.ExternalizedProperties); err != nil {
				return err
			}
		case "graphs":
			if err := json.Unmarshal([]byte(v), &strct.Graphs); err != nil {
				return err
			}
		case "invocations":
			if err := json.Unmarshal([]byte(v), &strct.Invocations); err != nil {
				return err
			}
		case "logicalLocations":
			if err := json.Unmarshal([]byte(v), &strct.LogicalLocations); err != nil {
				return err
			}
		case "policies":
			if err := json.Unmarshal([]byte(v), &strct.Policies); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "results":
			if err := json.Unmarshal([]byte(v), &strct.Results); err != nil {
				return err
			}
		case "taxonomies":
			if err := json.Unmarshal([]byte(v), &strct.Taxonomies); err != nil {
				return err
			}
		case "threadFlowLocations":
			if err := json.Unmarshal([]byte(v), &strct.ThreadFlowLocations); err != nil {
				return err
			}
		case "translations":
			if err := json.Unmarshal([]byte(v), &strct.Translations); err != nil {
				return err
			}
		case "webRequests":
			if err := json.Unmarshal([]byte(v), &strct.WebRequests); err != nil {
				return err
			}
		case "webResponses":
			if err := json.Unmarshal([]byte(v), &strct.WebResponses); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Fix) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ArtifactChanges" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "artifactChanges" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifactChanges\": ")
	if tmp, err := json.Marshal(strct.ArtifactChanges); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Fix) UnmarshalJSON(b []byte) error {
	artifactChangesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "artifactChanges":
			if err := json.Unmarshal([]byte(v), &strct.ArtifactChanges); err != nil {
				return err
			}
			artifactChangesReceived = true
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if artifactChanges (a required property) was received
	if !artifactChangesReceived {
		return errors.New("\"artifactChanges\" is required but was not present")
	}
	return nil
}

func (strct *Graph) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "edges" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"edges\": ")
	if tmp, err := json.Marshal(strct.Edges); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "nodes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"nodes\": ")
	if tmp, err := json.Marshal(strct.Nodes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Graph) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "edges":
			if err := json.Unmarshal([]byte(v), &strct.Edges); err != nil {
				return err
			}
		case "nodes":
			if err := json.Unmarshal([]byte(v), &strct.Nodes); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *GraphTraversal) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "edgeTraversals" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"edgeTraversals\": ")
	if tmp, err := json.Marshal(strct.EdgeTraversals); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "immutableState" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"immutableState\": ")
	if tmp, err := json.Marshal(strct.ImmutableState); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "initialState" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"initialState\": ")
	if tmp, err := json.Marshal(strct.InitialState); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "resultGraphIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resultGraphIndex\": ")
	if tmp, err := json.Marshal(strct.ResultGraphIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "runGraphIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"runGraphIndex\": ")
	if tmp, err := json.Marshal(strct.RunGraphIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GraphTraversal) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "edgeTraversals":
			if err := json.Unmarshal([]byte(v), &strct.EdgeTraversals); err != nil {
				return err
			}
		case "immutableState":
			if err := json.Unmarshal([]byte(v), &strct.ImmutableState); err != nil {
				return err
			}
		case "initialState":
			if err := json.Unmarshal([]byte(v), &strct.InitialState); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "resultGraphIndex":
			if err := json.Unmarshal([]byte(v), &strct.ResultGraphIndex); err != nil {
				return err
			}
		case "runGraphIndex":
			if err := json.Unmarshal([]byte(v), &strct.RunGraphIndex); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Invocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "account" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"account\": ")
	if tmp, err := json.Marshal(strct.Account); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "arguments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"arguments\": ")
	if tmp, err := json.Marshal(strct.Arguments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "commandLine" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"commandLine\": ")
	if tmp, err := json.Marshal(strct.CommandLine); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "endTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endTimeUtc\": ")
	if tmp, err := json.Marshal(strct.EndTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "environmentVariables" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"environmentVariables\": ")
	if tmp, err := json.Marshal(strct.EnvironmentVariables); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "executableLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"executableLocation\": ")
	if tmp, err := json.Marshal(strct.ExecutableLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ExecutionSuccessful" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "executionSuccessful" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"executionSuccessful\": ")
	if tmp, err := json.Marshal(strct.ExecutionSuccessful); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "exitCode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"exitCode\": ")
	if tmp, err := json.Marshal(strct.ExitCode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "exitCodeDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"exitCodeDescription\": ")
	if tmp, err := json.Marshal(strct.ExitCodeDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "exitSignalName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"exitSignalName\": ")
	if tmp, err := json.Marshal(strct.ExitSignalName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "exitSignalNumber" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"exitSignalNumber\": ")
	if tmp, err := json.Marshal(strct.ExitSignalNumber); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "machine" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"machine\": ")
	if tmp, err := json.Marshal(strct.Machine); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "notificationConfigurationOverrides" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"notificationConfigurationOverrides\": ")
	if tmp, err := json.Marshal(strct.NotificationConfigurationOverrides); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "processId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"processId\": ")
	if tmp, err := json.Marshal(strct.ProcessId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "processStartFailureMessage" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"processStartFailureMessage\": ")
	if tmp, err := json.Marshal(strct.ProcessStartFailureMessage); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "responseFiles" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"responseFiles\": ")
	if tmp, err := json.Marshal(strct.ResponseFiles); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ruleConfigurationOverrides" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ruleConfigurationOverrides\": ")
	if tmp, err := json.Marshal(strct.RuleConfigurationOverrides); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "startTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startTimeUtc\": ")
	if tmp, err := json.Marshal(strct.StartTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stderr" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stderr\": ")
	if tmp, err := json.Marshal(strct.Stderr); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stdin" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stdin\": ")
	if tmp, err := json.Marshal(strct.Stdin); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stdout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stdout\": ")
	if tmp, err := json.Marshal(strct.Stdout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stdoutStderr" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stdoutStderr\": ")
	if tmp, err := json.Marshal(strct.StdoutStderr); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "toolConfigurationNotifications" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"toolConfigurationNotifications\": ")
	if tmp, err := json.Marshal(strct.ToolConfigurationNotifications); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "toolExecutionNotifications" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"toolExecutionNotifications\": ")
	if tmp, err := json.Marshal(strct.ToolExecutionNotifications); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "workingDirectory" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"workingDirectory\": ")
	if tmp, err := json.Marshal(strct.WorkingDirectory); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Invocation) UnmarshalJSON(b []byte) error {
	executionSuccessfulReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "account":
			if err := json.Unmarshal([]byte(v), &strct.Account); err != nil {
				return err
			}
		case "arguments":
			if err := json.Unmarshal([]byte(v), &strct.Arguments); err != nil {
				return err
			}
		case "commandLine":
			if err := json.Unmarshal([]byte(v), &strct.CommandLine); err != nil {
				return err
			}
		case "endTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.EndTimeUtc); err != nil {
				return err
			}
		case "environmentVariables":
			if err := json.Unmarshal([]byte(v), &strct.EnvironmentVariables); err != nil {
				return err
			}
		case "executableLocation":
			if err := json.Unmarshal([]byte(v), &strct.ExecutableLocation); err != nil {
				return err
			}
		case "executionSuccessful":
			if err := json.Unmarshal([]byte(v), &strct.ExecutionSuccessful); err != nil {
				return err
			}
			executionSuccessfulReceived = true
		case "exitCode":
			if err := json.Unmarshal([]byte(v), &strct.ExitCode); err != nil {
				return err
			}
		case "exitCodeDescription":
			if err := json.Unmarshal([]byte(v), &strct.ExitCodeDescription); err != nil {
				return err
			}
		case "exitSignalName":
			if err := json.Unmarshal([]byte(v), &strct.ExitSignalName); err != nil {
				return err
			}
		case "exitSignalNumber":
			if err := json.Unmarshal([]byte(v), &strct.ExitSignalNumber); err != nil {
				return err
			}
		case "machine":
			if err := json.Unmarshal([]byte(v), &strct.Machine); err != nil {
				return err
			}
		case "notificationConfigurationOverrides":
			if err := json.Unmarshal([]byte(v), &strct.NotificationConfigurationOverrides); err != nil {
				return err
			}
		case "processId":
			if err := json.Unmarshal([]byte(v), &strct.ProcessId); err != nil {
				return err
			}
		case "processStartFailureMessage":
			if err := json.Unmarshal([]byte(v), &strct.ProcessStartFailureMessage); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "responseFiles":
			if err := json.Unmarshal([]byte(v), &strct.ResponseFiles); err != nil {
				return err
			}
		case "ruleConfigurationOverrides":
			if err := json.Unmarshal([]byte(v), &strct.RuleConfigurationOverrides); err != nil {
				return err
			}
		case "startTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.StartTimeUtc); err != nil {
				return err
			}
		case "stderr":
			if err := json.Unmarshal([]byte(v), &strct.Stderr); err != nil {
				return err
			}
		case "stdin":
			if err := json.Unmarshal([]byte(v), &strct.Stdin); err != nil {
				return err
			}
		case "stdout":
			if err := json.Unmarshal([]byte(v), &strct.Stdout); err != nil {
				return err
			}
		case "stdoutStderr":
			if err := json.Unmarshal([]byte(v), &strct.StdoutStderr); err != nil {
				return err
			}
		case "toolConfigurationNotifications":
			if err := json.Unmarshal([]byte(v), &strct.ToolConfigurationNotifications); err != nil {
				return err
			}
		case "toolExecutionNotifications":
			if err := json.Unmarshal([]byte(v), &strct.ToolExecutionNotifications); err != nil {
				return err
			}
		case "workingDirectory":
			if err := json.Unmarshal([]byte(v), &strct.WorkingDirectory); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if executionSuccessful (a required property) was received
	if !executionSuccessfulReceived {
		return errors.New("\"executionSuccessful\" is required but was not present")
	}
	return nil
}

func (strct *Location) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "annotations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"annotations\": ")
	if tmp, err := json.Marshal(strct.Annotations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "logicalLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"logicalLocations\": ")
	if tmp, err := json.Marshal(strct.LogicalLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "physicalLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"physicalLocation\": ")
	if tmp, err := json.Marshal(strct.PhysicalLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "relationships" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"relationships\": ")
	if tmp, err := json.Marshal(strct.Relationships); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Location) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "annotations":
			if err := json.Unmarshal([]byte(v), &strct.Annotations); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
		case "logicalLocations":
			if err := json.Unmarshal([]byte(v), &strct.LogicalLocations); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "physicalLocation":
			if err := json.Unmarshal([]byte(v), &strct.PhysicalLocation); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "relationships":
			if err := json.Unmarshal([]byte(v), &strct.Relationships); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *LocationRelationship) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kinds" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kinds\": ")
	if tmp, err := json.Marshal(strct.Kinds); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Target" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "target" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"target\": ")
	if tmp, err := json.Marshal(strct.Target); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LocationRelationship) UnmarshalJSON(b []byte) error {
	targetReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "kinds":
			if err := json.Unmarshal([]byte(v), &strct.Kinds); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "target":
			if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
				return err
			}
			targetReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if target (a required property) was received
	if !targetReceived {
		return errors.New("\"target\" is required but was not present")
	}
	return nil
}

func (strct *LogicalLocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "decoratedName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"decoratedName\": ")
	if tmp, err := json.Marshal(strct.DecoratedName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullyQualifiedName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullyQualifiedName\": ")
	if tmp, err := json.Marshal(strct.FullyQualifiedName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parentIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parentIndex\": ")
	if tmp, err := json.Marshal(strct.ParentIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *LogicalLocation) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "decoratedName":
			if err := json.Unmarshal([]byte(v), &strct.DecoratedName); err != nil {
				return err
			}
		case "fullyQualifiedName":
			if err := json.Unmarshal([]byte(v), &strct.FullyQualifiedName); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "kind":
			if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		case "parentIndex":
			if err := json.Unmarshal([]byte(v), &strct.ParentIndex); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Message) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "arguments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"arguments\": ")
	if tmp, err := json.Marshal(strct.Arguments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "markdown" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"markdown\": ")
	if tmp, err := json.Marshal(strct.Markdown); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "text" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"text\": ")
	if tmp, err := json.Marshal(strct.Text); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Message) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "arguments":
			if err := json.Unmarshal([]byte(v), &strct.Arguments); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
		case "markdown":
			if err := json.Unmarshal([]byte(v), &strct.Markdown); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "text":
			if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *MultiformatMessageString) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "markdown" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"markdown\": ")
	if tmp, err := json.Marshal(strct.Markdown); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Text" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "text" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"text\": ")
	if tmp, err := json.Marshal(strct.Text); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *MultiformatMessageString) UnmarshalJSON(b []byte) error {
	textReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "markdown":
			if err := json.Unmarshal([]byte(v), &strct.Markdown); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "text":
			if err := json.Unmarshal([]byte(v), &strct.Text); err != nil {
				return err
			}
			textReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if text (a required property) was received
	if !textReceived {
		return errors.New("\"text\" is required but was not present")
	}
	return nil
}

func (strct *Node) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "children" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"children\": ")
	if tmp, err := json.Marshal(strct.Children); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Id" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "label" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"label\": ")
	if tmp, err := json.Marshal(strct.Label); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Node) UnmarshalJSON(b []byte) error {
	idReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "children":
			if err := json.Unmarshal([]byte(v), &strct.Children); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
			idReceived = true
		case "label":
			if err := json.Unmarshal([]byte(v), &strct.Label); err != nil {
				return err
			}
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if id (a required property) was received
	if !idReceived {
		return errors.New("\"id\" is required but was not present")
	}
	return nil
}

func (strct *Notification) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "associatedRule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"associatedRule\": ")
	if tmp, err := json.Marshal(strct.AssociatedRule); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "descriptor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"descriptor\": ")
	if tmp, err := json.Marshal(strct.Descriptor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "exception" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"exception\": ")
	if tmp, err := json.Marshal(strct.Exception); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "level" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"level\": ")
	if tmp, err := json.Marshal(strct.Level); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "locations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"locations\": ")
	if tmp, err := json.Marshal(strct.Locations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Message" field is required
	if strct.Message == nil {
		return nil, errors.New("message is a required field")
	}
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "threadId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"threadId\": ")
	if tmp, err := json.Marshal(strct.ThreadId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "timeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timeUtc\": ")
	if tmp, err := json.Marshal(strct.TimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Notification) UnmarshalJSON(b []byte) error {
	messageReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "associatedRule":
			if err := json.Unmarshal([]byte(v), &strct.AssociatedRule); err != nil {
				return err
			}
		case "descriptor":
			if err := json.Unmarshal([]byte(v), &strct.Descriptor); err != nil {
				return err
			}
		case "exception":
			if err := json.Unmarshal([]byte(v), &strct.Exception); err != nil {
				return err
			}
		case "level":
			if err := json.Unmarshal([]byte(v), &strct.Level); err != nil {
				return err
			}
		case "locations":
			if err := json.Unmarshal([]byte(v), &strct.Locations); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
			messageReceived = true
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "threadId":
			if err := json.Unmarshal([]byte(v), &strct.ThreadId); err != nil {
				return err
			}
		case "timeUtc":
			if err := json.Unmarshal([]byte(v), &strct.TimeUtc); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if message (a required property) was received
	if !messageReceived {
		return errors.New("\"message\" is required but was not present")
	}
	return nil
}

func (strct *PhysicalLocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "address" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"address\": ")
	if tmp, err := json.Marshal(strct.Address); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "artifactLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifactLocation\": ")
	if tmp, err := json.Marshal(strct.ArtifactLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "contextRegion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"contextRegion\": ")
	if tmp, err := json.Marshal(strct.ContextRegion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "region" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"region\": ")
	if tmp, err := json.Marshal(strct.Region); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PhysicalLocation) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "address":
			if err := json.Unmarshal([]byte(v), &strct.Address); err != nil {
				return err
			}
		case "artifactLocation":
			if err := json.Unmarshal([]byte(v), &strct.ArtifactLocation); err != nil {
				return err
			}
		case "contextRegion":
			if err := json.Unmarshal([]byte(v), &strct.ContextRegion); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "region":
			if err := json.Unmarshal([]byte(v), &strct.Region); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *PropertyBag) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "tags" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tags\": ")
	if tmp, err := json.Marshal(strct.Tags); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal any additional Properties
	for k, v := range strct.AdditionalProperties {
		if comma {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", k))
		if tmp, err := json.Marshal(v); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		comma = true
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PropertyBag) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "tags":
			if err := json.Unmarshal([]byte(v), &strct.Tags); err != nil {
				return err
			}
		default:
			// an additional "interface{}" value
			var additionalValue interface{}
			if err := json.Unmarshal([]byte(v), &additionalValue); err != nil {
				return err // invalid additionalProperty
			}
			if strct.AdditionalProperties == nil {
				strct.AdditionalProperties = make(map[string]interface{}, 0)
			}
			strct.AdditionalProperties[k] = additionalValue
		}
	}
	return nil
}

func (strct *Rectangle) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "bottom" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"bottom\": ")
	if tmp, err := json.Marshal(strct.Bottom); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "left" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"left\": ")
	if tmp, err := json.Marshal(strct.Left); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "right" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"right\": ")
	if tmp, err := json.Marshal(strct.Right); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "top" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"top\": ")
	if tmp, err := json.Marshal(strct.Top); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Rectangle) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "bottom":
			if err := json.Unmarshal([]byte(v), &strct.Bottom); err != nil {
				return err
			}
		case "left":
			if err := json.Unmarshal([]byte(v), &strct.Left); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "right":
			if err := json.Unmarshal([]byte(v), &strct.Right); err != nil {
				return err
			}
		case "top":
			if err := json.Unmarshal([]byte(v), &strct.Top); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Region) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "byteLength" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"byteLength\": ")
	if tmp, err := json.Marshal(strct.ByteLength); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "byteOffset" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"byteOffset\": ")
	if tmp, err := json.Marshal(strct.ByteOffset); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "charLength" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"charLength\": ")
	if tmp, err := json.Marshal(strct.CharLength); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "charOffset" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"charOffset\": ")
	if tmp, err := json.Marshal(strct.CharOffset); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "endColumn" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endColumn\": ")
	if tmp, err := json.Marshal(strct.EndColumn); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "endLine" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endLine\": ")
	if tmp, err := json.Marshal(strct.EndLine); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "snippet" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"snippet\": ")
	if tmp, err := json.Marshal(strct.Snippet); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "sourceLanguage" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sourceLanguage\": ")
	if tmp, err := json.Marshal(strct.SourceLanguage); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "startColumn" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startColumn\": ")
	if tmp, err := json.Marshal(strct.StartColumn); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "startLine" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startLine\": ")
	if tmp, err := json.Marshal(strct.StartLine); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Region) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "byteLength":
			if err := json.Unmarshal([]byte(v), &strct.ByteLength); err != nil {
				return err
			}
		case "byteOffset":
			if err := json.Unmarshal([]byte(v), &strct.ByteOffset); err != nil {
				return err
			}
		case "charLength":
			if err := json.Unmarshal([]byte(v), &strct.CharLength); err != nil {
				return err
			}
		case "charOffset":
			if err := json.Unmarshal([]byte(v), &strct.CharOffset); err != nil {
				return err
			}
		case "endColumn":
			if err := json.Unmarshal([]byte(v), &strct.EndColumn); err != nil {
				return err
			}
		case "endLine":
			if err := json.Unmarshal([]byte(v), &strct.EndLine); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "snippet":
			if err := json.Unmarshal([]byte(v), &strct.Snippet); err != nil {
				return err
			}
		case "sourceLanguage":
			if err := json.Unmarshal([]byte(v), &strct.SourceLanguage); err != nil {
				return err
			}
		case "startColumn":
			if err := json.Unmarshal([]byte(v), &strct.StartColumn); err != nil {
				return err
			}
		case "startLine":
			if err := json.Unmarshal([]byte(v), &strct.StartLine); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Replacement) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "DeletedRegion" field is required
	if strct.DeletedRegion == nil {
		return nil, errors.New("deletedRegion is a required field")
	}
	// Marshal the "deletedRegion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"deletedRegion\": ")
	if tmp, err := json.Marshal(strct.DeletedRegion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "insertedContent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"insertedContent\": ")
	if tmp, err := json.Marshal(strct.InsertedContent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Replacement) UnmarshalJSON(b []byte) error {
	deletedRegionReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "deletedRegion":
			if err := json.Unmarshal([]byte(v), &strct.DeletedRegion); err != nil {
				return err
			}
			deletedRegionReceived = true
		case "insertedContent":
			if err := json.Unmarshal([]byte(v), &strct.InsertedContent); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if deletedRegion (a required property) was received
	if !deletedRegionReceived {
		return errors.New("\"deletedRegion\" is required but was not present")
	}
	return nil
}

func (strct *ReportingConfiguration) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "enabled" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"enabled\": ")
	if tmp, err := json.Marshal(strct.Enabled); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "level" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"level\": ")
	if tmp, err := json.Marshal(strct.Level); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parameters" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parameters\": ")
	if tmp, err := json.Marshal(strct.Parameters); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rank" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rank\": ")
	if tmp, err := json.Marshal(strct.Rank); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReportingConfiguration) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "enabled":
			if err := json.Unmarshal([]byte(v), &strct.Enabled); err != nil {
				return err
			}
		case "level":
			if err := json.Unmarshal([]byte(v), &strct.Level); err != nil {
				return err
			}
		case "parameters":
			if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "rank":
			if err := json.Unmarshal([]byte(v), &strct.Rank); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ReportingDescriptor) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "defaultConfiguration" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"defaultConfiguration\": ")
	if tmp, err := json.Marshal(strct.DefaultConfiguration); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "deprecatedGuids" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"deprecatedGuids\": ")
	if tmp, err := json.Marshal(strct.DeprecatedGuids); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "deprecatedIds" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"deprecatedIds\": ")
	if tmp, err := json.Marshal(strct.DeprecatedIds); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "deprecatedNames" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"deprecatedNames\": ")
	if tmp, err := json.Marshal(strct.DeprecatedNames); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullDescription\": ")
	if tmp, err := json.Marshal(strct.FullDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "help" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"help\": ")
	if tmp, err := json.Marshal(strct.Help); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "helpUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"helpUri\": ")
	if tmp, err := json.Marshal(strct.HelpUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Id" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "messageStrings" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"messageStrings\": ")
	if tmp, err := json.Marshal(strct.MessageStrings); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "relationships" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"relationships\": ")
	if tmp, err := json.Marshal(strct.Relationships); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "shortDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shortDescription\": ")
	if tmp, err := json.Marshal(strct.ShortDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReportingDescriptor) UnmarshalJSON(b []byte) error {
	idReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "defaultConfiguration":
			if err := json.Unmarshal([]byte(v), &strct.DefaultConfiguration); err != nil {
				return err
			}
		case "deprecatedGuids":
			if err := json.Unmarshal([]byte(v), &strct.DeprecatedGuids); err != nil {
				return err
			}
		case "deprecatedIds":
			if err := json.Unmarshal([]byte(v), &strct.DeprecatedIds); err != nil {
				return err
			}
		case "deprecatedNames":
			if err := json.Unmarshal([]byte(v), &strct.DeprecatedNames); err != nil {
				return err
			}
		case "fullDescription":
			if err := json.Unmarshal([]byte(v), &strct.FullDescription); err != nil {
				return err
			}
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "help":
			if err := json.Unmarshal([]byte(v), &strct.Help); err != nil {
				return err
			}
		case "helpUri":
			if err := json.Unmarshal([]byte(v), &strct.HelpUri); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
			idReceived = true
		case "messageStrings":
			if err := json.Unmarshal([]byte(v), &strct.MessageStrings); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "relationships":
			if err := json.Unmarshal([]byte(v), &strct.Relationships); err != nil {
				return err
			}
		case "shortDescription":
			if err := json.Unmarshal([]byte(v), &strct.ShortDescription); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if id (a required property) was received
	if !idReceived {
		return errors.New("\"id\" is required but was not present")
	}
	return nil
}

func (strct *ReportingDescriptorReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "toolComponent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"toolComponent\": ")
	if tmp, err := json.Marshal(strct.ToolComponent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReportingDescriptorReference) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "toolComponent":
			if err := json.Unmarshal([]byte(v), &strct.ToolComponent); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ReportingDescriptorRelationship) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kinds" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kinds\": ")
	if tmp, err := json.Marshal(strct.Kinds); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Target" field is required
	if strct.Target == nil {
		return nil, errors.New("target is a required field")
	}
	// Marshal the "target" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"target\": ")
	if tmp, err := json.Marshal(strct.Target); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ReportingDescriptorRelationship) UnmarshalJSON(b []byte) error {
	targetReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "kinds":
			if err := json.Unmarshal([]byte(v), &strct.Kinds); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "target":
			if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
				return err
			}
			targetReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if target (a required property) was received
	if !targetReceived {
		return errors.New("\"target\" is required but was not present")
	}
	return nil
}

func (strct *Result) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "analysisTarget" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"analysisTarget\": ")
	if tmp, err := json.Marshal(strct.AnalysisTarget); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "attachments" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"attachments\": ")
	if tmp, err := json.Marshal(strct.Attachments); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "baselineState" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"baselineState\": ")
	if tmp, err := json.Marshal(strct.BaselineState); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "codeFlows" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"codeFlows\": ")
	if tmp, err := json.Marshal(strct.CodeFlows); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "correlationGuid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"correlationGuid\": ")
	if tmp, err := json.Marshal(strct.CorrelationGuid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fingerprints" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fingerprints\": ")
	if tmp, err := json.Marshal(strct.Fingerprints); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fixes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fixes\": ")
	if tmp, err := json.Marshal(strct.Fixes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "graphTraversals" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"graphTraversals\": ")
	if tmp, err := json.Marshal(strct.GraphTraversals); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "graphs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"graphs\": ")
	if tmp, err := json.Marshal(strct.Graphs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "hostedViewerUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"hostedViewerUri\": ")
	if tmp, err := json.Marshal(strct.HostedViewerUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "level" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"level\": ")
	if tmp, err := json.Marshal(strct.Level); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "locations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"locations\": ")
	if tmp, err := json.Marshal(strct.Locations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Message" field is required
	if strct.Message == nil {
		return nil, errors.New("message is a required field")
	}
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "occurrenceCount" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"occurrenceCount\": ")
	if tmp, err := json.Marshal(strct.OccurrenceCount); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "partialFingerprints" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"partialFingerprints\": ")
	if tmp, err := json.Marshal(strct.PartialFingerprints); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "provenance" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"provenance\": ")
	if tmp, err := json.Marshal(strct.Provenance); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rank" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rank\": ")
	if tmp, err := json.Marshal(strct.Rank); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "relatedLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"relatedLocations\": ")
	if tmp, err := json.Marshal(strct.RelatedLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rule\": ")
	if tmp, err := json.Marshal(strct.Rule); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ruleId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ruleId\": ")
	if tmp, err := json.Marshal(strct.RuleId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ruleIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ruleIndex\": ")
	if tmp, err := json.Marshal(strct.RuleIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stacks" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stacks\": ")
	if tmp, err := json.Marshal(strct.Stacks); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "suppressions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"suppressions\": ")
	if tmp, err := json.Marshal(strct.Suppressions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "taxa" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"taxa\": ")
	if tmp, err := json.Marshal(strct.Taxa); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webRequest" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webRequest\": ")
	if tmp, err := json.Marshal(strct.WebRequest); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webResponse" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webResponse\": ")
	if tmp, err := json.Marshal(strct.WebResponse); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "workItemUris" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"workItemUris\": ")
	if tmp, err := json.Marshal(strct.WorkItemUris); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Result) UnmarshalJSON(b []byte) error {
	messageReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "analysisTarget":
			if err := json.Unmarshal([]byte(v), &strct.AnalysisTarget); err != nil {
				return err
			}
		case "attachments":
			if err := json.Unmarshal([]byte(v), &strct.Attachments); err != nil {
				return err
			}
		case "baselineState":
			if err := json.Unmarshal([]byte(v), &strct.BaselineState); err != nil {
				return err
			}
		case "codeFlows":
			if err := json.Unmarshal([]byte(v), &strct.CodeFlows); err != nil {
				return err
			}
		case "correlationGuid":
			if err := json.Unmarshal([]byte(v), &strct.CorrelationGuid); err != nil {
				return err
			}
		case "fingerprints":
			if err := json.Unmarshal([]byte(v), &strct.Fingerprints); err != nil {
				return err
			}
		case "fixes":
			if err := json.Unmarshal([]byte(v), &strct.Fixes); err != nil {
				return err
			}
		case "graphTraversals":
			if err := json.Unmarshal([]byte(v), &strct.GraphTraversals); err != nil {
				return err
			}
		case "graphs":
			if err := json.Unmarshal([]byte(v), &strct.Graphs); err != nil {
				return err
			}
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "hostedViewerUri":
			if err := json.Unmarshal([]byte(v), &strct.HostedViewerUri); err != nil {
				return err
			}
		case "kind":
			if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
				return err
			}
		case "level":
			if err := json.Unmarshal([]byte(v), &strct.Level); err != nil {
				return err
			}
		case "locations":
			if err := json.Unmarshal([]byte(v), &strct.Locations); err != nil {
				return err
			}
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
			messageReceived = true
		case "occurrenceCount":
			if err := json.Unmarshal([]byte(v), &strct.OccurrenceCount); err != nil {
				return err
			}
		case "partialFingerprints":
			if err := json.Unmarshal([]byte(v), &strct.PartialFingerprints); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "provenance":
			if err := json.Unmarshal([]byte(v), &strct.Provenance); err != nil {
				return err
			}
		case "rank":
			if err := json.Unmarshal([]byte(v), &strct.Rank); err != nil {
				return err
			}
		case "relatedLocations":
			if err := json.Unmarshal([]byte(v), &strct.RelatedLocations); err != nil {
				return err
			}
		case "rule":
			if err := json.Unmarshal([]byte(v), &strct.Rule); err != nil {
				return err
			}
		case "ruleId":
			if err := json.Unmarshal([]byte(v), &strct.RuleId); err != nil {
				return err
			}
		case "ruleIndex":
			if err := json.Unmarshal([]byte(v), &strct.RuleIndex); err != nil {
				return err
			}
		case "stacks":
			if err := json.Unmarshal([]byte(v), &strct.Stacks); err != nil {
				return err
			}
		case "suppressions":
			if err := json.Unmarshal([]byte(v), &strct.Suppressions); err != nil {
				return err
			}
		case "taxa":
			if err := json.Unmarshal([]byte(v), &strct.Taxa); err != nil {
				return err
			}
		case "webRequest":
			if err := json.Unmarshal([]byte(v), &strct.WebRequest); err != nil {
				return err
			}
		case "webResponse":
			if err := json.Unmarshal([]byte(v), &strct.WebResponse); err != nil {
				return err
			}
		case "workItemUris":
			if err := json.Unmarshal([]byte(v), &strct.WorkItemUris); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if message (a required property) was received
	if !messageReceived {
		return errors.New("\"message\" is required but was not present")
	}
	return nil
}

func (strct *ResultProvenance) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "conversionSources" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"conversionSources\": ")
	if tmp, err := json.Marshal(strct.ConversionSources); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "firstDetectionRunGuid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"firstDetectionRunGuid\": ")
	if tmp, err := json.Marshal(strct.FirstDetectionRunGuid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "firstDetectionTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"firstDetectionTimeUtc\": ")
	if tmp, err := json.Marshal(strct.FirstDetectionTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "invocationIndex" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"invocationIndex\": ")
	if tmp, err := json.Marshal(strct.InvocationIndex); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "lastDetectionRunGuid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"lastDetectionRunGuid\": ")
	if tmp, err := json.Marshal(strct.LastDetectionRunGuid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "lastDetectionTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"lastDetectionTimeUtc\": ")
	if tmp, err := json.Marshal(strct.LastDetectionTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ResultProvenance) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "conversionSources":
			if err := json.Unmarshal([]byte(v), &strct.ConversionSources); err != nil {
				return err
			}
		case "firstDetectionRunGuid":
			if err := json.Unmarshal([]byte(v), &strct.FirstDetectionRunGuid); err != nil {
				return err
			}
		case "firstDetectionTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.FirstDetectionTimeUtc); err != nil {
				return err
			}
		case "invocationIndex":
			if err := json.Unmarshal([]byte(v), &strct.InvocationIndex); err != nil {
				return err
			}
		case "lastDetectionRunGuid":
			if err := json.Unmarshal([]byte(v), &strct.LastDetectionRunGuid); err != nil {
				return err
			}
		case "lastDetectionTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.LastDetectionTimeUtc); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Run) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "addresses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"addresses\": ")
	if tmp, err := json.Marshal(strct.Addresses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "artifacts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"artifacts\": ")
	if tmp, err := json.Marshal(strct.Artifacts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "automationDetails" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"automationDetails\": ")
	if tmp, err := json.Marshal(strct.AutomationDetails); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "baselineGuid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"baselineGuid\": ")
	if tmp, err := json.Marshal(strct.BaselineGuid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "columnKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"columnKind\": ")
	if tmp, err := json.Marshal(strct.ColumnKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "conversion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"conversion\": ")
	if tmp, err := json.Marshal(strct.Conversion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "defaultEncoding" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"defaultEncoding\": ")
	if tmp, err := json.Marshal(strct.DefaultEncoding); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "defaultSourceLanguage" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"defaultSourceLanguage\": ")
	if tmp, err := json.Marshal(strct.DefaultSourceLanguage); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "externalPropertyFileReferences" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"externalPropertyFileReferences\": ")
	if tmp, err := json.Marshal(strct.ExternalPropertyFileReferences); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "graphs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"graphs\": ")
	if tmp, err := json.Marshal(strct.Graphs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "invocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"invocations\": ")
	if tmp, err := json.Marshal(strct.Invocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "language" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"language\": ")
	if tmp, err := json.Marshal(strct.Language); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "logicalLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"logicalLocations\": ")
	if tmp, err := json.Marshal(strct.LogicalLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "newlineSequences" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"newlineSequences\": ")
	if tmp, err := json.Marshal(strct.NewlineSequences); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "originalUriBaseIds" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"originalUriBaseIds\": ")
	if tmp, err := json.Marshal(strct.OriginalUriBaseIds); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "policies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"policies\": ")
	if tmp, err := json.Marshal(strct.Policies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "redactionTokens" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"redactionTokens\": ")
	if tmp, err := json.Marshal(strct.RedactionTokens); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "results" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"results\": ")
	if tmp, err := json.Marshal(strct.Results); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "runAggregates" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"runAggregates\": ")
	if tmp, err := json.Marshal(strct.RunAggregates); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "specialLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"specialLocations\": ")
	if tmp, err := json.Marshal(strct.SpecialLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "taxonomies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"taxonomies\": ")
	if tmp, err := json.Marshal(strct.Taxonomies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "threadFlowLocations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"threadFlowLocations\": ")
	if tmp, err := json.Marshal(strct.ThreadFlowLocations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Tool" field is required
	if strct.Tool == nil {
		return nil, errors.New("tool is a required field")
	}
	// Marshal the "tool" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tool\": ")
	if tmp, err := json.Marshal(strct.Tool); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "translations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"translations\": ")
	if tmp, err := json.Marshal(strct.Translations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "versionControlProvenance" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"versionControlProvenance\": ")
	if tmp, err := json.Marshal(strct.VersionControlProvenance); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webRequests" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webRequests\": ")
	if tmp, err := json.Marshal(strct.WebRequests); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webResponses" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webResponses\": ")
	if tmp, err := json.Marshal(strct.WebResponses); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Run) UnmarshalJSON(b []byte) error {
	toolReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "addresses":
			if err := json.Unmarshal([]byte(v), &strct.Addresses); err != nil {
				return err
			}
		case "artifacts":
			if err := json.Unmarshal([]byte(v), &strct.Artifacts); err != nil {
				return err
			}
		case "automationDetails":
			if err := json.Unmarshal([]byte(v), &strct.AutomationDetails); err != nil {
				return err
			}
		case "baselineGuid":
			if err := json.Unmarshal([]byte(v), &strct.BaselineGuid); err != nil {
				return err
			}
		case "columnKind":
			if err := json.Unmarshal([]byte(v), &strct.ColumnKind); err != nil {
				return err
			}
		case "conversion":
			if err := json.Unmarshal([]byte(v), &strct.Conversion); err != nil {
				return err
			}
		case "defaultEncoding":
			if err := json.Unmarshal([]byte(v), &strct.DefaultEncoding); err != nil {
				return err
			}
		case "defaultSourceLanguage":
			if err := json.Unmarshal([]byte(v), &strct.DefaultSourceLanguage); err != nil {
				return err
			}
		case "externalPropertyFileReferences":
			if err := json.Unmarshal([]byte(v), &strct.ExternalPropertyFileReferences); err != nil {
				return err
			}
		case "graphs":
			if err := json.Unmarshal([]byte(v), &strct.Graphs); err != nil {
				return err
			}
		case "invocations":
			if err := json.Unmarshal([]byte(v), &strct.Invocations); err != nil {
				return err
			}
		case "language":
			if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
				return err
			}
		case "logicalLocations":
			if err := json.Unmarshal([]byte(v), &strct.LogicalLocations); err != nil {
				return err
			}
		case "newlineSequences":
			if err := json.Unmarshal([]byte(v), &strct.NewlineSequences); err != nil {
				return err
			}
		case "originalUriBaseIds":
			if err := json.Unmarshal([]byte(v), &strct.OriginalUriBaseIds); err != nil {
				return err
			}
		case "policies":
			if err := json.Unmarshal([]byte(v), &strct.Policies); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "redactionTokens":
			if err := json.Unmarshal([]byte(v), &strct.RedactionTokens); err != nil {
				return err
			}
		case "results":
			if err := json.Unmarshal([]byte(v), &strct.Results); err != nil {
				return err
			}
		case "runAggregates":
			if err := json.Unmarshal([]byte(v), &strct.RunAggregates); err != nil {
				return err
			}
		case "specialLocations":
			if err := json.Unmarshal([]byte(v), &strct.SpecialLocations); err != nil {
				return err
			}
		case "taxonomies":
			if err := json.Unmarshal([]byte(v), &strct.Taxonomies); err != nil {
				return err
			}
		case "threadFlowLocations":
			if err := json.Unmarshal([]byte(v), &strct.ThreadFlowLocations); err != nil {
				return err
			}
		case "tool":
			if err := json.Unmarshal([]byte(v), &strct.Tool); err != nil {
				return err
			}
			toolReceived = true
		case "translations":
			if err := json.Unmarshal([]byte(v), &strct.Translations); err != nil {
				return err
			}
		case "versionControlProvenance":
			if err := json.Unmarshal([]byte(v), &strct.VersionControlProvenance); err != nil {
				return err
			}
		case "webRequests":
			if err := json.Unmarshal([]byte(v), &strct.WebRequests); err != nil {
				return err
			}
		case "webResponses":
			if err := json.Unmarshal([]byte(v), &strct.WebResponses); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if tool (a required property) was received
	if !toolReceived {
		return errors.New("\"tool\" is required but was not present")
	}
	return nil
}

func (strct *RunAutomationDetails) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "correlationGuid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"correlationGuid\": ")
	if tmp, err := json.Marshal(strct.CorrelationGuid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "description" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"description\": ")
	if tmp, err := json.Marshal(strct.Description); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RunAutomationDetails) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "correlationGuid":
			if err := json.Unmarshal([]byte(v), &strct.CorrelationGuid); err != nil {
				return err
			}
		case "description":
			if err := json.Unmarshal([]byte(v), &strct.Description); err != nil {
				return err
			}
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *SpecialLocations) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "displayBase" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"displayBase\": ")
	if tmp, err := json.Marshal(strct.DisplayBase); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SpecialLocations) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "displayBase":
			if err := json.Unmarshal([]byte(v), &strct.DisplayBase); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Stack) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Frames" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "frames" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"frames\": ")
	if tmp, err := json.Marshal(strct.Frames); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Stack) UnmarshalJSON(b []byte) error {
	framesReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "frames":
			if err := json.Unmarshal([]byte(v), &strct.Frames); err != nil {
				return err
			}
			framesReceived = true
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if frames (a required property) was received
	if !framesReceived {
		return errors.New("\"frames\" is required but was not present")
	}
	return nil
}

func (strct *StackFrame) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "module" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"module\": ")
	if tmp, err := json.Marshal(strct.Module); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parameters" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parameters\": ")
	if tmp, err := json.Marshal(strct.Parameters); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "threadId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"threadId\": ")
	if tmp, err := json.Marshal(strct.ThreadId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *StackFrame) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
		case "module":
			if err := json.Unmarshal([]byte(v), &strct.Module); err != nil {
				return err
			}
		case "parameters":
			if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "threadId":
			if err := json.Unmarshal([]byte(v), &strct.ThreadId); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *SARIF) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "inlineExternalProperties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"inlineExternalProperties\": ")
	if tmp, err := json.Marshal(strct.InlineExternalProperties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Runs" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "runs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"runs\": ")
	if tmp, err := json.Marshal(strct.Runs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "$schema" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"$schema\": ")
	if tmp, err := json.Marshal(strct.Schema); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Version" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SARIF) UnmarshalJSON(b []byte) error {
	runsReceived := false
	versionReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "inlineExternalProperties":
			if err := json.Unmarshal([]byte(v), &strct.InlineExternalProperties); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "runs":
			if err := json.Unmarshal([]byte(v), &strct.Runs); err != nil {
				return err
			}
			runsReceived = true
		case "$schema":
			if err := json.Unmarshal([]byte(v), &strct.Schema); err != nil {
				return err
			}
		case "version":
			if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
				return err
			}
			versionReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if runs (a required property) was received
	if !runsReceived {
		return errors.New("\"runs\" is required but was not present")
	}
	// check if version (a required property) was received
	if !versionReceived {
		return errors.New("\"version\" is required but was not present")
	}
	return nil
}

func (strct *Suppression) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "justification" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"justification\": ")
	if tmp, err := json.Marshal(strct.Justification); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Kind" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "kind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "state" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"state\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Suppression) UnmarshalJSON(b []byte) error {
	kindReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "justification":
			if err := json.Unmarshal([]byte(v), &strct.Justification); err != nil {
				return err
			}
		case "kind":
			if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
				return err
			}
			kindReceived = true
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "state":
			if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if kind (a required property) was received
	if !kindReceived {
		return errors.New("\"kind\" is required but was not present")
	}
	return nil
}

func (strct *ThreadFlow) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "id" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "immutableState" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"immutableState\": ")
	if tmp, err := json.Marshal(strct.ImmutableState); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "initialState" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"initialState\": ")
	if tmp, err := json.Marshal(strct.InitialState); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Locations" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "locations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"locations\": ")
	if tmp, err := json.Marshal(strct.Locations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "message" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"message\": ")
	if tmp, err := json.Marshal(strct.Message); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ThreadFlow) UnmarshalJSON(b []byte) error {
	locationsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "id":
			if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
				return err
			}
		case "immutableState":
			if err := json.Unmarshal([]byte(v), &strct.ImmutableState); err != nil {
				return err
			}
		case "initialState":
			if err := json.Unmarshal([]byte(v), &strct.InitialState); err != nil {
				return err
			}
		case "locations":
			if err := json.Unmarshal([]byte(v), &strct.Locations); err != nil {
				return err
			}
			locationsReceived = true
		case "message":
			if err := json.Unmarshal([]byte(v), &strct.Message); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if locations (a required property) was received
	if !locationsReceived {
		return errors.New("\"locations\" is required but was not present")
	}
	return nil
}

func (strct *ThreadFlowLocation) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "executionOrder" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"executionOrder\": ")
	if tmp, err := json.Marshal(strct.ExecutionOrder); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "executionTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"executionTimeUtc\": ")
	if tmp, err := json.Marshal(strct.ExecutionTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "importance" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"importance\": ")
	if tmp, err := json.Marshal(strct.Importance); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "kinds" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kinds\": ")
	if tmp, err := json.Marshal(strct.Kinds); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "location" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"location\": ")
	if tmp, err := json.Marshal(strct.Location); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "module" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"module\": ")
	if tmp, err := json.Marshal(strct.Module); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "nestingLevel" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"nestingLevel\": ")
	if tmp, err := json.Marshal(strct.NestingLevel); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stack" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stack\": ")
	if tmp, err := json.Marshal(strct.Stack); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "state" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"state\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "taxa" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"taxa\": ")
	if tmp, err := json.Marshal(strct.Taxa); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webRequest" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webRequest\": ")
	if tmp, err := json.Marshal(strct.WebRequest); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "webResponse" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"webResponse\": ")
	if tmp, err := json.Marshal(strct.WebResponse); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ThreadFlowLocation) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "executionOrder":
			if err := json.Unmarshal([]byte(v), &strct.ExecutionOrder); err != nil {
				return err
			}
		case "executionTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.ExecutionTimeUtc); err != nil {
				return err
			}
		case "importance":
			if err := json.Unmarshal([]byte(v), &strct.Importance); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "kinds":
			if err := json.Unmarshal([]byte(v), &strct.Kinds); err != nil {
				return err
			}
		case "location":
			if err := json.Unmarshal([]byte(v), &strct.Location); err != nil {
				return err
			}
		case "module":
			if err := json.Unmarshal([]byte(v), &strct.Module); err != nil {
				return err
			}
		case "nestingLevel":
			if err := json.Unmarshal([]byte(v), &strct.NestingLevel); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "stack":
			if err := json.Unmarshal([]byte(v), &strct.Stack); err != nil {
				return err
			}
		case "state":
			if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
				return err
			}
		case "taxa":
			if err := json.Unmarshal([]byte(v), &strct.Taxa); err != nil {
				return err
			}
		case "webRequest":
			if err := json.Unmarshal([]byte(v), &strct.WebRequest); err != nil {
				return err
			}
		case "webResponse":
			if err := json.Unmarshal([]byte(v), &strct.WebResponse); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Tool) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Driver" field is required
	if strct.Driver == nil {
		return nil, errors.New("driver is a required field")
	}
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "extensions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extensions\": ")
	if tmp, err := json.Marshal(strct.Extensions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Tool) UnmarshalJSON(b []byte) error {
	driverReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
			driverReceived = true
		case "extensions":
			if err := json.Unmarshal([]byte(v), &strct.Extensions); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if driver (a required property) was received
	if !driverReceived {
		return errors.New("\"driver\" is required but was not present")
	}
	return nil
}

func (strct *ToolComponent) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "associatedComponent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"associatedComponent\": ")
	if tmp, err := json.Marshal(strct.AssociatedComponent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "contents" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"contents\": ")
	if tmp, err := json.Marshal(strct.Contents); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "dottedQuadFileVersion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dottedQuadFileVersion\": ")
	if tmp, err := json.Marshal(strct.DottedQuadFileVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "downloadUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"downloadUri\": ")
	if tmp, err := json.Marshal(strct.DownloadUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullDescription\": ")
	if tmp, err := json.Marshal(strct.FullDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullName\": ")
	if tmp, err := json.Marshal(strct.FullName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "globalMessageStrings" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"globalMessageStrings\": ")
	if tmp, err := json.Marshal(strct.GlobalMessageStrings); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "informationUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"informationUri\": ")
	if tmp, err := json.Marshal(strct.InformationUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "isComprehensive" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"isComprehensive\": ")
	if tmp, err := json.Marshal(strct.IsComprehensive); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "language" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"language\": ")
	if tmp, err := json.Marshal(strct.Language); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "localizedDataSemanticVersion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"localizedDataSemanticVersion\": ")
	if tmp, err := json.Marshal(strct.LocalizedDataSemanticVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "locations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"locations\": ")
	if tmp, err := json.Marshal(strct.Locations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "minimumRequiredLocalizedDataSemanticVersion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"minimumRequiredLocalizedDataSemanticVersion\": ")
	if tmp, err := json.Marshal(strct.MinimumRequiredLocalizedDataSemanticVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "notifications" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"notifications\": ")
	if tmp, err := json.Marshal(strct.Notifications); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "organization" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"organization\": ")
	if tmp, err := json.Marshal(strct.Organization); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "product" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"product\": ")
	if tmp, err := json.Marshal(strct.Product); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "productSuite" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"productSuite\": ")
	if tmp, err := json.Marshal(strct.ProductSuite); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "releaseDateUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"releaseDateUtc\": ")
	if tmp, err := json.Marshal(strct.ReleaseDateUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rules" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rules\": ")
	if tmp, err := json.Marshal(strct.Rules); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "semanticVersion" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"semanticVersion\": ")
	if tmp, err := json.Marshal(strct.SemanticVersion); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "shortDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shortDescription\": ")
	if tmp, err := json.Marshal(strct.ShortDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "supportedTaxonomies" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"supportedTaxonomies\": ")
	if tmp, err := json.Marshal(strct.SupportedTaxonomies); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "taxa" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"taxa\": ")
	if tmp, err := json.Marshal(strct.Taxa); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "translationMetadata" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"translationMetadata\": ")
	if tmp, err := json.Marshal(strct.TranslationMetadata); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ToolComponent) UnmarshalJSON(b []byte) error {
	nameReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "associatedComponent":
			if err := json.Unmarshal([]byte(v), &strct.AssociatedComponent); err != nil {
				return err
			}
		case "contents":
			if err := json.Unmarshal([]byte(v), &strct.Contents); err != nil {
				return err
			}
		case "dottedQuadFileVersion":
			if err := json.Unmarshal([]byte(v), &strct.DottedQuadFileVersion); err != nil {
				return err
			}
		case "downloadUri":
			if err := json.Unmarshal([]byte(v), &strct.DownloadUri); err != nil {
				return err
			}
		case "fullDescription":
			if err := json.Unmarshal([]byte(v), &strct.FullDescription); err != nil {
				return err
			}
		case "fullName":
			if err := json.Unmarshal([]byte(v), &strct.FullName); err != nil {
				return err
			}
		case "globalMessageStrings":
			if err := json.Unmarshal([]byte(v), &strct.GlobalMessageStrings); err != nil {
				return err
			}
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "informationUri":
			if err := json.Unmarshal([]byte(v), &strct.InformationUri); err != nil {
				return err
			}
		case "isComprehensive":
			if err := json.Unmarshal([]byte(v), &strct.IsComprehensive); err != nil {
				return err
			}
		case "language":
			if err := json.Unmarshal([]byte(v), &strct.Language); err != nil {
				return err
			}
		case "localizedDataSemanticVersion":
			if err := json.Unmarshal([]byte(v), &strct.LocalizedDataSemanticVersion); err != nil {
				return err
			}
		case "locations":
			if err := json.Unmarshal([]byte(v), &strct.Locations); err != nil {
				return err
			}
		case "minimumRequiredLocalizedDataSemanticVersion":
			if err := json.Unmarshal([]byte(v), &strct.MinimumRequiredLocalizedDataSemanticVersion); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "notifications":
			if err := json.Unmarshal([]byte(v), &strct.Notifications); err != nil {
				return err
			}
		case "organization":
			if err := json.Unmarshal([]byte(v), &strct.Organization); err != nil {
				return err
			}
		case "product":
			if err := json.Unmarshal([]byte(v), &strct.Product); err != nil {
				return err
			}
		case "productSuite":
			if err := json.Unmarshal([]byte(v), &strct.ProductSuite); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "releaseDateUtc":
			if err := json.Unmarshal([]byte(v), &strct.ReleaseDateUtc); err != nil {
				return err
			}
		case "rules":
			if err := json.Unmarshal([]byte(v), &strct.Rules); err != nil {
				return err
			}
		case "semanticVersion":
			if err := json.Unmarshal([]byte(v), &strct.SemanticVersion); err != nil {
				return err
			}
		case "shortDescription":
			if err := json.Unmarshal([]byte(v), &strct.ShortDescription); err != nil {
				return err
			}
		case "supportedTaxonomies":
			if err := json.Unmarshal([]byte(v), &strct.SupportedTaxonomies); err != nil {
				return err
			}
		case "taxa":
			if err := json.Unmarshal([]byte(v), &strct.Taxa); err != nil {
				return err
			}
		case "translationMetadata":
			if err := json.Unmarshal([]byte(v), &strct.TranslationMetadata); err != nil {
				return err
			}
		case "version":
			if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	return nil
}

func (strct *ToolComponentReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "guid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"guid\": ")
	if tmp, err := json.Marshal(strct.Guid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ToolComponentReference) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "guid":
			if err := json.Unmarshal([]byte(v), &strct.Guid); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *TranslationMetadata) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "downloadUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"downloadUri\": ")
	if tmp, err := json.Marshal(strct.DownloadUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullDescription\": ")
	if tmp, err := json.Marshal(strct.FullDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "fullName" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"fullName\": ")
	if tmp, err := json.Marshal(strct.FullName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "informationUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"informationUri\": ")
	if tmp, err := json.Marshal(strct.InformationUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "shortDescription" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shortDescription\": ")
	if tmp, err := json.Marshal(strct.ShortDescription); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TranslationMetadata) UnmarshalJSON(b []byte) error {
	nameReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "downloadUri":
			if err := json.Unmarshal([]byte(v), &strct.DownloadUri); err != nil {
				return err
			}
		case "fullDescription":
			if err := json.Unmarshal([]byte(v), &strct.FullDescription); err != nil {
				return err
			}
		case "fullName":
			if err := json.Unmarshal([]byte(v), &strct.FullName); err != nil {
				return err
			}
		case "informationUri":
			if err := json.Unmarshal([]byte(v), &strct.InformationUri); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "shortDescription":
			if err := json.Unmarshal([]byte(v), &strct.ShortDescription); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	return nil
}

func (strct *VersionControlDetails) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "asOfTimeUtc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"asOfTimeUtc\": ")
	if tmp, err := json.Marshal(strct.AsOfTimeUtc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "branch" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"branch\": ")
	if tmp, err := json.Marshal(strct.Branch); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "mappedTo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"mappedTo\": ")
	if tmp, err := json.Marshal(strct.MappedTo); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "RepositoryUri" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "repositoryUri" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"repositoryUri\": ")
	if tmp, err := json.Marshal(strct.RepositoryUri); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "revisionId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"revisionId\": ")
	if tmp, err := json.Marshal(strct.RevisionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "revisionTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"revisionTag\": ")
	if tmp, err := json.Marshal(strct.RevisionTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *VersionControlDetails) UnmarshalJSON(b []byte) error {
	repositoryUriReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "asOfTimeUtc":
			if err := json.Unmarshal([]byte(v), &strct.AsOfTimeUtc); err != nil {
				return err
			}
		case "branch":
			if err := json.Unmarshal([]byte(v), &strct.Branch); err != nil {
				return err
			}
		case "mappedTo":
			if err := json.Unmarshal([]byte(v), &strct.MappedTo); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "repositoryUri":
			if err := json.Unmarshal([]byte(v), &strct.RepositoryUri); err != nil {
				return err
			}
			repositoryUriReceived = true
		case "revisionId":
			if err := json.Unmarshal([]byte(v), &strct.RevisionId); err != nil {
				return err
			}
		case "revisionTag":
			if err := json.Unmarshal([]byte(v), &strct.RevisionTag); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if repositoryUri (a required property) was received
	if !repositoryUriReceived {
		return errors.New("\"repositoryUri\" is required but was not present")
	}
	return nil
}

func (strct *WebRequest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "body" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"body\": ")
	if tmp, err := json.Marshal(strct.Body); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "headers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"headers\": ")
	if tmp, err := json.Marshal(strct.Headers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "method" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"method\": ")
	if tmp, err := json.Marshal(strct.Method); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parameters" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parameters\": ")
	if tmp, err := json.Marshal(strct.Parameters); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "protocol" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"protocol\": ")
	if tmp, err := json.Marshal(strct.Protocol); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "target" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"target\": ")
	if tmp, err := json.Marshal(strct.Target); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *WebRequest) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "body":
			if err := json.Unmarshal([]byte(v), &strct.Body); err != nil {
				return err
			}
		case "headers":
			if err := json.Unmarshal([]byte(v), &strct.Headers); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "method":
			if err := json.Unmarshal([]byte(v), &strct.Method); err != nil {
				return err
			}
		case "parameters":
			if err := json.Unmarshal([]byte(v), &strct.Parameters); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "protocol":
			if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
				return err
			}
		case "target":
			if err := json.Unmarshal([]byte(v), &strct.Target); err != nil {
				return err
			}
		case "version":
			if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *WebResponse) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "body" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"body\": ")
	if tmp, err := json.Marshal(strct.Body); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "headers" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"headers\": ")
	if tmp, err := json.Marshal(strct.Headers); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "index" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"index\": ")
	if tmp, err := json.Marshal(strct.Index); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "noResponseReceived" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"noResponseReceived\": ")
	if tmp, err := json.Marshal(strct.NoResponseReceived); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "properties" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"properties\": ")
	if tmp, err := json.Marshal(strct.Properties); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "protocol" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"protocol\": ")
	if tmp, err := json.Marshal(strct.Protocol); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "reasonPhrase" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"reasonPhrase\": ")
	if tmp, err := json.Marshal(strct.ReasonPhrase); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "statusCode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"statusCode\": ")
	if tmp, err := json.Marshal(strct.StatusCode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *WebResponse) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "body":
			if err := json.Unmarshal([]byte(v), &strct.Body); err != nil {
				return err
			}
		case "headers":
			if err := json.Unmarshal([]byte(v), &strct.Headers); err != nil {
				return err
			}
		case "index":
			if err := json.Unmarshal([]byte(v), &strct.Index); err != nil {
				return err
			}
		case "noResponseReceived":
			if err := json.Unmarshal([]byte(v), &strct.NoResponseReceived); err != nil {
				return err
			}
		case "properties":
			if err := json.Unmarshal([]byte(v), &strct.Properties); err != nil {
				return err
			}
		case "protocol":
			if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
				return err
			}
		case "reasonPhrase":
			if err := json.Unmarshal([]byte(v), &strct.ReasonPhrase); err != nil {
				return err
			}
		case "statusCode":
			if err := json.Unmarshal([]byte(v), &strct.StatusCode); err != nil {
				return err
			}
		case "version":
			if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}
