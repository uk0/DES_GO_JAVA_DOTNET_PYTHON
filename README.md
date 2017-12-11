# DES_GO_JAVA_DOTNET_PYTHON
DES 加密数据通用 四种语言的加密解密 。


* Python

```bash
➜  expbox python demo.py
    53c8927a4dd1e9b6f44c9d53af92c794
    Hello World!    
```

* Golang

```bash
go run *.go

# 结果
53c8927a4dd1e9b6f44c9d53af92c794
Hello World!

[Done] exited with code=0 in 0.776 seconds
```

* C#(.NET)

```bash
dotnet run Program.cs

```

* Program.cs

```C#
using System;
using System.Security.Cryptography;
using System.IO; 
using System.Text;
namespace myApp
{
    public class Program
    {
         static void Main(string[] args)
        {
             Console.WriteLine("Hello World!");
                    string data = "Hello World!";
            string encryptData =  DES.Encrypt(data,"Passw0rd","Passw0rd");
            Console.WriteLine("加密后数据 1："+encryptData);

            string decryptData =  DES.Decrypt(encryptData,"Passw0rd","Passw0rd");
             Console.WriteLine("解密后数据 1："+decryptData);
        }
    }
}
```


* Java

```java
   //EncryptHelper 内
   public static void main(String[] args) throws Exception {
        System.out.println(EncryptHelper.desDecrypt("53c8927a4dd1e9b6f44c9d53af92c794"));
        System.out.println(EncryptHelper.desEncrypt("Hello World!"));

    }

```

```bash
Hello World!
53C8927A4DD1E9B6F44C9D53AF92C794

Process finished with exit code 0
```

