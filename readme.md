# Sort Algorithms in Golang

This repository have implementations of sorting methods using Golang. If you want to run the benchmark tests, you can run:
```
make bench
```

## Simple Sort
The complexity is calculated like this:

For $N = 3$ we have:
$$(2+1+0) = 3 = 3*2/2 = 3$$

For $N=4$ we have:
$$3+2+1+0 = 6 = 4*3/2 = 6$$

For $N=13$ we have:
$$12+11+10+9+8+7+6+5+4+3+2+1+0 = 78 = 13*12/2 = 78$$

So the complexity is

$$ O(N-1 + N-2 + N-3 + N-4 ...) = O(N*(N-1)/2) $$

With simplification we have: $O(Nˆ2)$

We don't need more memory usage to make this process so, for the space complexity we have: $O(1)$


## Other Sort Algorithms

<table>
    <tr>
        <th>Algorithms</th>
        <th>Best Time complexity</th>
        <th>Average Time complexity</th>
        <th>Worst Time complexity</th>
        <th>Space complexity</th>
    </tr>
    <tr>
        <td>Simple Sort</td>
        <td>O(Nˆ2)</td>
        <td>O(N^2)</td>
        <td>O(Nˆ2)</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Bubble Sort</td>
        <td>O(N)</td>
        <td>O(N^2)</td>
        <td>O(Nˆ2)</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Selection Sort</td>
        <td>O(Nˆ2)</td>
        <td>O(N^2)</td>
        <td>O(Nˆ2)</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Insertion Sort</td>
        <td>O(N)</td>
        <td>O(N^2)</td>
        <td>O(Nˆ2)</td>
        <td>O(1)</td>
    </tr>
    <tr>
        <td>Merge Sort</td>
        <td>O(N*logN)</td>
        <td>O(N*logN)</td>
        <td>O(N*logN)</td>
        <td>O(N)</td>
    </tr>
    <!-- <tr>
        <td>Quick Sort</td>
        <td>O(N*logN)</td>
        <td>O(N*logN)</td>
        <td>O(N^2)</td>
        <td>O(N)</td>
    </tr>
    <tr>
        <td>Heap Sort</td>
        <td>O(N*logN)</td>
        <td>O(N*logN)</td>
        <td>O(N*logN)</td>
        <td>O(1)</td>
    </tr> -->
</table>