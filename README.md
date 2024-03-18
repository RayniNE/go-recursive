# Merge Sort Algorythm

Merge Sort is an algorythm that's used to order array of data. It is another method of the Bubble Sort/Quick Sort family, but in this case, it eliminates the nested loops, providing a faster sort but also making it more stable. By stable i mean, that if the same number is repeated, the number will be ordered alongside the duplicate.

But it has a drawback, it uses recursion and it creates a second copy of the original array, making it more memory intensive than its counterpart.

To run the app, clone the repo and run in the cmd:

```
go run main.go
```

Example of the result:

```
go run main.go                                                                                                                                        ─╯
Unsorted Array
[114 99 99 69 7 186 103 108 177 170 9 123 113 21 159 9 195 112 4 55 112 157 14 93 18 199 45 185 171 189 130 20 89 55 0 53 39 91 185 18 54 21 54 97 86 25 93 20 54 110 131 85 187 22 31 8 93 67 155 49 152 9 135 33 102 5 26 193 135 118 35 181 10 143 40 188 56 92 104 122 153 47 167 174 128 194 106 22 89 36 66 140 169 157 20 137 92 89 26 40]


Sorted Array
[0 4 5 7 8 9 9 9 10 14 18 18 20 20 20 21 21 22 22 25 26 26 31 33 35 36 39 40 40 45 47 49 53 54 54 54 55 55 56 66 67 69 85 86 89 89 89 91 92 92 93 93 93 97 99 99 102 103 104 106 108 110 112 112 113 114 118 122 123 128 130 131 135 135 137 140 143 152 153 155 157 157 159 167 169 170 171 174 177 181 185 185 186 187 188 189 193 194 195 199]
```
