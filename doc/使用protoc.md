## protoc常用命令
- 查看帮助
```
protoc --help


Usage: protoc [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
                              standard input and write the raw tag/value
                              pairs in text format to standard output.  No
                              PROTO_FILES should be given when using this
                              flag.
  --descriptor_set_in=FILES   Specifies a delimited list of FILES
                              each containing a FileDescriptorSet (a
                              protocol buffer defined in descriptor.proto).
                              The FileDescriptor for each of the PROTO_FILES
                              provided will be loaded from these
                              FileDescriptorSets. If a FileDescriptor
                              appears multiple times, the first occurrence
                              will be used.
  -oFILE,                     Writes a FileDescriptorSet (a protocol buffer,
    --descriptor_set_out=FILE defined in descriptor.proto) containing all of
                              the input files to FILE.
  --include_imports           When using --descriptor_set_out, also include
                              all dependencies of the input files in the
                              set, so that the set is self-contained.
  --include_source_info       When using --descriptor_set_out, do not strip
                              SourceCodeInfo from the FileDescriptorProto.
                              This results in vastly larger descriptors that
                              include information about the original
                              location of each decl in the source file as
                              well as surrounding comments.
  --dependency_out=FILE       Write a dependency output file in the format
                              expected by make. This writes the transitive
                              set of input file paths to FILE
  --error_format=FORMAT       Set the format in which to print errors.
                              FORMAT may be 'gcc' (the default) or 'msvs'
                              (Microsoft Visual Studio format).
  --print_free_field_numbers  Print the free field numbers of the messages
                              defined in the given proto files. Groups share
                              the same field number space with the parent
                              message. Extension ranges are counted as
                              occupied fields numbers.

  --plugin=EXECUTABLE         Specifies a plugin executable to use.
                              Normally, protoc searches the PATH for
                              plugins, but you may specify additional
                              executables not in the path using this flag.
                              Additionally, EXECUTABLE may be of the form
                              NAME=PATH, in which case the given plugin name
                              is mapped to the given executable even if
                              the executable's own name differs.
  --cpp_out=OUT_DIR           Generate C++ header and source.
  --csharp_out=OUT_DIR        Generate C# source file.
  --java_out=OUT_DIR          Generate Java source file.
  --js_out=OUT_DIR            Generate JavaScript source.
  --objc_out=OUT_DIR          Generate Objective C header and source.
  --php_out=OUT_DIR           Generate PHP source file.
  --python_out=OUT_DIR        Generate Python source file.
  --ruby_out=OUT_DIR          Generate Ruby source file.
  @<filename>                 Read options and filenames from file. If a
                              relative file path is specified, the file
                              will be searched in the working directory.
                              The --proto_path option will not affect how
                              this argument file is searched. Content of
                              the file will be expanded in the position of
                              @<filename> as in the argument list. Note
                              that shell expansion is not applied to the
                              content of the file (i.e., you cannot use
                              quotes, wildcards, escapes, commands, etc.).
                              Each line corresponds to a single argument,
                              even if it contains spaces.
```

## 用法解析
```
Usage: protoc [OPTION] PROTO_FILES
用法: protoc 可选项 proto 文件

Parse PROTO_FILES and generate output based on the options given:
解析PROTO_FILES并根据给出的选项生成输出:
```
```
```
- 导入路径
```
-IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
指定要在其中搜索导入的目录。可以多次指定;
将按顺序搜索目录。如果没有指定，则使用当前工作目录。 
                              
```


- 查看protoc版本
```
--version                   Show version info and exit.
显示版本信息并退出。

示例: protoc --version
libprotoc 3.6.1
```
- 查看帮助信息
```
-h, --help                  Show this text and exit.
显示帮助文本并退出

示例: 
protoc --help
protoc -h
```
- 编码 从文本到二进制
```
--encode=MESSAGE_TYPE       Read a text-format message of the given type
                          from standard input and write it in binary
                          to standard output.  The message type must
                          be defined in PROTO_FILES or their imports.
                              
从标准输入读取给定类型的文本格式消息，并将其以二进制形式写入标准输出。
消息类型必须在 proto_files 或其导入中定义。                              
```
- 解码 从二进制到文本
```
--decode=MESSAGE_TYPE       Read a binary message of the given type from
                          standard input and write it in text format
                          to standard output.  The message type must
                          be defined in PROTO_FILES or their imports.
从标准输入读取给定类型的二进制消息，并将其以文本格式写入标准输出。
消息类型必须在proto_files或其导入中定义。                              
```
- 
```
--decode_raw                Read an arbitrary protocol message from
                          standard input and write the raw tag/value
                          pairs in text format to standard output.  No
                          PROTO_FILES should be given when using this
                          flag.
从标准输入读取任意协议消息，并将原始的tag/value对以文本格式写入标准输出。
使用此标志时不应提供proto_files。                          
```
- 
```
--descriptor_set_in=FILES   Specifies a delimited list of FILES
                          each containing a FileDescriptorSet (a
                          protocol buffer defined in descriptor.proto).
                          The FileDescriptor for each of the PROTO_FILES
                          provided will be loaded from these
                          FileDescriptorSets. If a FileDescriptor
                          appears multiple times, the first occurrence
                          will be used.
指定包含FileDescriptorSet(在description .proto中定义的协议缓冲区)的文件分隔列表。
提供的每个PROTO_FILES的文件描述符将从这些文件描述符集中加载。
如果文件描述符出现多次，将使用第一次出现。                          
```