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

#### *结论*
从数据结果中我们可以得出，java在字符串根据正则表达式进行字符替换的时候，性能表现突出，其他的场景均不如golang的表现。而我们通过动态链接库的方式，替换golang的正则引擎为rust的时候，性能表现很不均衡，而golang原生的sdk 正则引擎则有均衡的表现。因此，无特殊指定优化场景的情况下，建议还是使用golang原生的sdk引擎。其他语言的引擎有待继续验证。
