# Async communication protocol

## Important

**One event payload should be produced only by one service!**

**Consumers should make decidions about event handling only based on payload type!**

## Structure

```
<root>
  |- proto
     |- context_name
	    |- service_package
```

This package is a schema registry for recognizable and pluggable event communication.
+ Event producing
+ Event consuming
+ Topic naming definitions
+ Extending protocol

