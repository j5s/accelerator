# Accelerator

## Practice

### log4j-core-2.14.0.jar

Rule
```text
INVOKEINTERFACE javax/naming/Context.lookup *
```

Result
```text
org/apache/logging/log4j/core/net/JndiManager lookup
```

### spring-cloud-gateway-server-3.0.6.jar

Rule
```text
INVOKEINTERFACE org/springframework/expression/Expression.getValue *
```

Result
```text
org/springframework/cloud/gateway/discovery/DiscoveryClientRouteDefinitionLocator buildRouteDefinition
org/springframework/cloud/gateway/discovery/DiscoveryClientRouteDefinitionLocator getValueFromExpr
org/springframework/cloud/gateway/discovery/DiscoveryClientRouteDefinitionLocator lambda$getRouteDefinitions$2
org/springframework/cloud/gateway/support/ShortcutConfigurable getValue
```

Rule
```text
INVOKESPECIAL org/springframework/expression/spel/support/StandardEvaluationContext.<init> *
INVOKEVIRTUAL org/springframework/expression/spel/standard/SpelExpressionParser.parseExpression *
INVOKEINTERFACE org/springframework/expression/Expression.getValue *
```

Result
```text
org/springframework/cloud/gateway/support/ShortcutConfigurable getValue
```

### spring-cloud-function-context-3.1.0.jar

Rule
```text
INVOKEINTERFACE org/springframework/expression/Expression.getValue *
```

Result
```text
org/springframework/cloud/function/context/catalog/SimpleFunctionRegistry$FunctionInvocationWrapper parseMultipleValueArguments
org/springframework/cloud/function/context/config/RoutingFunction functionFromExpression
```

Rule
```text
INVOKESPECIAL org/springframework/expression/spel/support/StandardEvaluationContext.<init> *
```

Result
```text
org/springframework/cloud/function/context/config/RoutingFunction <init>
```

## reference

[Java Class File Format](https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html)

[JVM Instruction Set](https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html)