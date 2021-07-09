this project is main to compare about golang & java performance of regular expression

#### backgroud
调用次数： 10000次
正则表达式：
```
findImgTag:<img[^<>]*\ssrc=(?:'|")?([^\s'"<>]+)(?:'|")?[^<>]*>
findAllTag: <[^<>]*>
findWhiteSpace:[\s　]*
```

#### result
| engine | replace | findImgTag | findAllTag | findWhiteSpace |
| ---- | ---- | ---- | ---- | ---- |
| java01 | 7758ms | 3253ms| 1683ms | 3395ms |
| java02 | 7462ms | 3258ms | 1675ms | 3403ms |
| sdk01 | 17634.334ms| 239.567ms| 39.969ms | 48.095ms |
| sdk02 | 17912.696ms | 245.925ms | 38.734ms | 43.421ms |
| rust01 | 17595.447ms | 51.610ms | 429.589ms | 18275.605ms |
| rust02 | 17342.340ms | 50.287ms | 389.774 | 18080.263ms |
