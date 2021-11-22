import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';

if(process.env.NODE_ENV === "development") {
    console.log("Axios API Mock running");
    const mock = new MockAdapter(axios);

    mock.onGet(/\/api\/main\/problem.*/).reply((req) => {
        console.log("[Mock:Req] " + req.url);
        const body = {
        error: null,
        content: {
            title: "색칠 1",
            link: "https://www.acmicpc.net/problem/1117",
            description: `<p>지민이는 종이에 색칠하기를 좋아한다. 지민이는 W×H 크기의 직사각형 종이를 가지고 있다. 지민이는 종이에 다음과 같이 색칠 하려고 한다.</p><ol><li>종이를 x = f에 맞춰서 접는다. 이때, 왼쪽 종이가 오른쪽 종이 위에 올라오게 접는다.</li><li>종이를 가로로 c+1개의 크기가 동일 한 구간으로 나눈다. 그 다음에 c번 가장 위의 구간부터 차례대로 접는다.</li><li>왼쪽 아래가 (x<sub>1</sub>, y<sub>1</sub>) 이고, 오른쪽 위가 (x<sub>2</sub>, y<sub>2</sub>)인 직사각형을 찾는다. 이때, (0, 0)은 현재 접힌 상태에서 가장 왼쪽 아래 점이다. 그 직사각형을 칠한다. 이때, 페인트는 겹쳐있는 모든 곳에 스며든다.</li><li>종이를 편다.</li></ol><p>다음 예는 5×6 종이에, x=2이고, c=2이고, (x<sub>1</sub>, y<sub>1</sub>) = (1, 1), (x<sub>2</sub>, y<sub>2</sub>) = (3, 2)인 경우이다.</p><p style="text-align: center;"><img alt="" src="https://upload.acmicpc.net/7c49e41f-720b-4add-ad4c-76c14713a041/-/preview/"/><img alt="" src="https://upload.acmicpc.net/a7f8f028-ed9b-4453-a94a-eacb09e26377/-/preview/"/><img alt="" src="https://upload.acmicpc.net/3c67a1ba-49a4-4d1b-a5aa-e0147144adab/-/preview/"/></p><p style="text-align: center;"><img alt="" src="https://upload.acmicpc.net/534a806e-29a9-4045-ad48-c0da65b59412/-/preview/"/><img alt="" src="https://upload.acmicpc.net/a8ccf82d-3d34-45d0-b71a-85e9b945eea5/-/preview/"/><img alt="" src="https://upload.acmicpc.net/3abf2512-9a2a-40a1-9465-17d87c7cd58c/-/preview/"/></p><p>W, H, f, c, x<sub>1</sub>, y<sub>1</sub>, x<sub>2</sub>, y<sub>2</sub>가 주어질 때, 색칠되어 있지 않은 면적을 구하는 프로그램을 작성하시오.</p>`,
        },
        };
        return [200, body];
    });

    mock.onGet(/\/api\/main\/discussion.*/).reply((req) => {
        console.log("[Mock:Req]" + req.url);
        const body = {
        error: null,
        content: {
            alt: "https://www.acmicpc.net/status?from_problem=1&problem_id=1117",
            links: [],
        },
        };
        return [200, body];
    });

    mock.onGet(/\/api\/difficulty\/problem.*/).reply((req) => {
        console.log("[Mock:Req] " + req.url);
        const body = {
            error: null,
            content: {
                title: "450. Delete Node in a BST",
                link: "https://leetcode.com/problems/delete-node-in-a-bst/",
                description: `<div><p>Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.</p>

<p>Basically, the deletion can be divided into two stages:</p>

<ol>
	<li>Search for a node to remove.</li>
	<li>If the node is found, delete the node.</li>
</ol>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg" style="width: 800px; height: 214px;">
<pre><strong>Input:</strong> root = [5,3,6,2,4,null,7], key = 3
<strong>Output:</strong> [5,4,6,2,null,null,7]
<strong>Explanation:</strong> Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
<img alt="" src="https://assets.leetcode.com/uploads/2020/09/04/del_node_supp.jpg" style="width: 350px; height: 255px;">
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> root = [5,3,6,2,4,null,7], key = 0
<strong>Output:</strong> [5,3,6,2,4,null,7]
<strong>Explanation:</strong> The tree does not contain a node with value = 0.
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> root = [], key = 0
<strong>Output:</strong> []
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li>The number of nodes in the tree is in the range <code>[0, 10<sup>4</sup>]</code>.</li>
	<li><code>-10<sup>5</sup> &lt;= Node.val &lt;= 10<sup>5</sup></code></li>
	<li>Each node has a <strong>unique</strong> value.</li>
	<li><code>root</code> is a valid binary search tree.</li>
	<li><code>-10<sup>5</sup> &lt;= key &lt;= 10<sup>5</sup></code></li>
</ul>

<p>&nbsp;</p>
<p><strong>Follow up:</strong> Could you solve it with time complexity <code>O(height of tree)</code>?</p>
</div>`,
            },
        };

        return [200, body];
    });

    mock.onGet(/\/api\/difficulty\/discussion.*/).reply((req) => {
        console.log("[Mock:Req]" + req.url);
        const body = {
        error: null,
        content: {
            alt: "",
            links: [
                { title: "✔️[C++] Shortest Recursive Solution | Detailed Explanation with images", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590789/C%2B%2B-Shortest-Recursive-Solution-or-Detailed-Explanation-with-images", },
                { title: "✅ [C++] Simple Solution w/ Images & Detailed Explanation | Iterative & Recursive Approach", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591176/C%2B%2B-Simple-Solution-w-Images-and-Detailed-Explanation-or-Iterative-and-Recursive-Approach", },
                { title: "✔ JAVA | Recursive | Most Intutive | Proper Explanation Using Image", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590868/JAVA-or-Recursive-or-Most-Intutive-or-Proper-Explanation-Using-Image", },
                { title: "C++ Short and Simple Recursive Solution, Detailed Explanation", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590757/C%2B%2B-Short-and-Simple-Recursive-Solution-Detailed-Explanation", },
                { title: "Python solution without predecessor and successor", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591048/Python-solution-without-predecessor-and-successor", },
                { title: "C++ || Recursive || Iterative || Detailed and simple solution", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590768/C%2B%2B-oror-Recursive-oror-Iterative-oror-Detailed-and-simple-solution", },
                { title: "[C++] Recursive Solution", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590935/C%2B%2B-Recursive-Solution", },
                { title: "[C++] Explained Solution || Recursive Approach", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591316/C%2B%2B-Explained-Solution-oror-Recursive-Approach", },
                { title: "JAVA --> 100 % faster || recursive", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590970/JAVA-greater-100-faster-oror-recursive", },
                { title: "c clean", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591416/c-clean", },
                { title: "C++ with comments", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591393/C%2B%2B-with-comments", },
                { title: "C++ || 90% Faster Recursive Solution || Easy to Understand", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591309/C%2B%2B-oror-90-Faster-Recursive-Solution-oror-Easy-to-Understand", },
                { title: "Recursive sol with all conditions explained through diagrams & commented code in O(H)time", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591208/Recursive-sol-with-all-conditions-explained-through-diagrams-and-commented-code-in-O(H)time", },
                { title: "[Python] Tree height complexity, iterative, explained", link: "https://leetcode.com/problems/delete-node-in-a-bst/discuss/1590983/Python-Tree-height-complexity-iterative-explained", },
            ],
        },
        };
        return [200, body];
    });
}
