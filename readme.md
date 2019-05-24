# Bindex
General purpose: store entity IDs in several keys. Make bit operations on different keys and get result IDs  

> Input slice must contains **unique** values! For example, entity IDs.
  Otherwise index may returns wrong result.  

### Usage example 
```
index := New()

index.Set("KEY1", []int{10, 20, 30, 40})
index.Set("KEY2", []int{11, 20, 30, 50})

result := index.Select("KEY1").And("KEY2").Result()

// Result contains []int{20, 30}
```

### Support operations
- And
- AndNot
- Or

For more examples see bindex_test file