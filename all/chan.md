**команды** [`go-ipfs-cmds`](https://github.com/ipfs/go-ipfs-cmds) [![Трэвис CI](https://camo.githubusercontent.com/65b349d04aaa4d685764816a75c6f28f869d2ad0/68747470733a2f2f7472617669732d63692e636f6d2f697066732f676f2d697066732d636d64732e7376673f6272616e63683d6d6173746572)](https://travis-ci.com/ipfs/go-ipfs-cmds) [![codecov](https://camo.githubusercontent.com/afee9c54558e0d22b2ca6326bb0ad139e276e2b9/68747470733a2f2f636f6465636f762e696f2f67682f697066732f676f2d697066732d636d64732f6272616e63682f6d61737465722f67726170682f62616467652e737667)](https://codecov.io/gh/ipfs/go-ipfs-cmds)  Библиотека команд CLI и HTTP  [`go-ipfs-api`](https://github.com/ipfs/go-ipfs-api) [![Трэвис CI](https://camo.githubusercontent.com/df4da5b26743efd7c0a8e82df35be10b4bcfe440/68747470733a2f2f7472617669732d63692e636f6d2f697066732f676f2d697066732d6170692e7376673f6272616e63683d6d6173746572)](https://travis-ci.com/ipfs/go-ipfs-api) [![codecov](https://camo.githubusercontent.com/6dda58a220a376b379ab6c14070fba2c673afc27/68747470733a2f2f636f6465636f762e696f2f67682f697066732f676f2d697066732d6170692f6272616e63682f6d61737465722f67726170682f62616467652e737667)](https://codecov.io/gh/ipfs/go-ipfs-api)  оболочка для API IPFS HTTP  **Метрики и логирование**


```golang
// Check chanel function
func ChainCall(fns ...func() (*Chain, error)) (err error) {
    for _, fn := range fns {
        if _, err = fn(); err != nil {
            break
        }
    }
    return
}
```
