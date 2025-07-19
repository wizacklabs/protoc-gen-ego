# protc-gen-go

The protoc-gen-go project is another protoc plugin to generate Go code for both proto2 and proto3 versions of the protocol buffer language. 

The protoc-gen-go works just like official utility, and provides some new features to auto generate message field tag and custimize field name from comments.

For more information about the usage of this plugin, see: https://protobuf.dev/reference/go/go-generated.

## Features
- auto generate message field tags from comments
- customize message field name

## Limitations
- only works for message field declaration

## Installation
```bash
go install gitee.com/wizacklabs/proto-gen-go
```

## Usage
- auto generate field tags
    
    protobuf source code
    ```protobuf
    message foo {
        // @gorm.tag=column:id;autoIncrement
        // @json.tag=ID
        int64 id = 1;
    }
    ```
    generated go source code
    ```go
    type Foo struct {
        state         protoimpl.MessageState
        sizeCache     protoimpl.SizeCache
        unknownFields protoimpl.UnknownFields

        Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"ID,omitempty" gorm:"column:id;autoIncrement"`
    }
    ```

- customize field name

    protobuf source code
    ```protobuf
    message foo {
        // @go.name=ID
        int64 id = 1;
    }
    ```
    generated go source code
    ```go
    type Foo struct {
        state         protoimpl.MessageState
        sizeCache     protoimpl.SizeCache
        unknownFields protoimpl.UnknownFields

        ID int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
    }
    ```
