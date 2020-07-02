// Copyright 2020 SunJun <i@sjis.me>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package algorithms

// 莱温斯坦距离: https://www.cnblogs.com/danvid/p/10877231.html
// 1.如果 i=0 或者 j=0，edit(i,j)=j 或者i
// 2.如果 i&&j>=1 则edit(i,j) = min{edit(i-1,j)+1, edit(i,j-1)+1, edit(i-1,j-1)+x} 如果m[i]=n[j]相等，则x=0，否则x=1；
// 其中 edit(i,j) 表示字符串 m,n  前i,j 位字符的编辑距离
func LevenshteinDistance(m, n []rune) int {

	i := len(m)
	j := len(n)

	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}

	min := LevenshteinDistance(m[:i-1], n) + 1

	if m := LevenshteinDistance(m, n[:j-1]) + 1; m < min {
		min = m
	}

	x := 0
	if m[i-1] != n[j-1] {
		x = 1
	}
	if m := LevenshteinDistance(m[:i-1], n[:j-1]) + x; m < min {
		min = m
	}
	return min
}
