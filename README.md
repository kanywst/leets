# EASY

## TwoSum 11/1

```
整数 nums と整数 target の配列が与えられた場合，2 つの数値のインデックスを返し，それらの足し算が target になるようにします．

各入力が正確に1つの解を持つと仮定して、同じ要素を2回使用することはできません。
```

* HashMap利用する

```python
def twoSum(nums, target):
    hashmap = {}
    for i, num in enumerate(nums):
        search_key = target - num
        if search_key not in hashmap:
            hashmap[num] = i
        else:
            return [i, hashmap[search_key]]
```