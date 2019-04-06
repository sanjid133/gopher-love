# gopher-love
Send love to all your fellow go library by giving Github  *

### Installation

To install run:
```console
go get github.com/sanjid133/gopher-love/...

$ gopher -help
```

### Configuration

To config run following command

```console
$ gopher config
```

To generate [github](github.com) token follow [this](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line).


## Love

- To love all the repository of a user or organization, run:

    ```console
    $ gopher love -o github.com/<org/user name>
    ```

- To love all the repository of a list of users or organizations:

    - Create a text file containing the link of the users or organizations as below:
        ```
        github.com/kubernetes
        github.com/golang
        github.com/sanjid133
        ``` 
    - Then run:
        ```console
        $ gopher love -f <path-of-the-created-file>
        ```

- To love current dependency, run

    ```console
    $ gopher love
    ```

## Supported Dependency

 - [Dep](https://golang.github.io/dep/docs/introduction.html)
 - [Glide](https://github.com/Masterminds/glide)
 
## Supported Platform

 - [Github](github.com)
  

