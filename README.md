# Accelerator

Batch analysis jar packages using Golang

Use to detect security vulnerabilities

This will be much simpler and faster than Java ASM

## Quick Start

accelerator need a rule file (default: rule.txt)

Enter a directory of jar files, accelerator will scan all jar files and extract them

```shell
./accelerator -rule your_rule_file -jars your_jar_dir
```

In the writing of rule, only **INVOKE** instruction is supported at present

```text
INVOKEVIRTUAL ... *
```

If a single instruction is written, the detection is successful if the corresponding instruction in a method

Usually, the desc attribute is not easy to remember, so wildcards such as * are supported

```text
INVOKEVIRTUAL [first rule] *
INVOKEVIRTUAL [next rule] *
...
```

Rules that support multiple **INVOKE** instructions at the same time

If the instruction set in the target method matches the calling order of multiple instructions, it is considered to match

## How it Works

(1) Unzip the jar file to get all the class files

(2) The class file is parsed according to the Oracle Java Specification

(3) Parse all methods in the method area of all classes to obtain the instruction set

(4) Improve instruction content by finding constant pool

(5) Parse the user rule and match it with the current method instruction set

## Examples

Native SQL Inject Rule
```text
INVOKEVIRTUAL java/lang/StringBuilder.append *
INVOKEINTERFACE java/sql/Statement.executeQuery *
```

SQL Inject JdbcTemplate Rule
```text
INVOKEVIRTUAL java/lang/StringBuilder.append *
INVOKEVIRTUAL org/springframework/jdbc/core/JdbcTemplate.query *
```

Simple RCE Rule
```text
INVOKEVIRTUAL java/lang/Runtime.exec *
```

Simple RCE Rule (Command Inject)
```text
INVOKEVIRTUAL java/lang/StringBuilder.append *
INVOKEVIRTUAL java/lang/Runtime.exec *
```

Some SSRF Rule
- INVOKEVIRTUAL java/net/URL.openConnection *
- INVOKEVIRTUAL org/apache/http/impl/client/CloseableHttpClient.execute *
- INVOKEINTERFACE okhttp3/Call.execute *

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