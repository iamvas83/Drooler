Implementation of Binary Trees
1. Binary Tree Program in C
``` C
//structure that contains data, address of the left child, and the address of the right child
struct node {int item;
 struct node* left;struct node* right;
};
 
// function to create a new node
struct node* createNode(int data) {
 
 //allocating space for new nodestruct node *node = (struct node *)malloc(sizeof(struct node));
 
 //inserting data in the node
 node->item = data;
  //setting left and right child as NULL
 node->left = NULL;
 node->right = NULL;
 
 return node;
}
 
// Insert to the left of the node
struct node* insertLeft(struct node* root, int value) {
 root->left = createNode(value);
 return root->left;
}
 
// Inserting to the right of the node
struct node* insertRight(struct node* root, int value) {
 root->right = createNode(value);
 return root->right;
}

```

``` c++
2. Binary Tree Program in C++
// Binary Tree in C++
//structure that contains data, address of left child, address of the right child
struct Node {int data;
 struct node *left;struct node *right;
};
 
// function to create a new node
Node *newNode(int data) {
 //allocating space for the node
 Node *node = new Node;
 
 //storing in the data
 node->data = data;
  //setting left and right children to NULL
 node->left = NULL;
 node->right = NULL;
 return (node);
}

```

``` java
3. Binary Tree Program in Java
// Binary Tree in Java
 
// creating a node that holds the data, address of the left child, and the address of the right child
class Node {
 int key;
 Node left, right;
 
 //setting data in the nodepublic Node(int item) {
   key = item;
   //setting left and right child equal to NULL
   left = right = null;
 }


class BinaryTree {
 Node root;
  //inserting data into the binary tree
 BinaryTree(int key) {
   root = new Node(key);
 }
 
 //set root NULL when the binary tree is created for the first time
 BinaryTree() {
   root = null;
 }
 
 public static void main(String[] args) {
 
   //creating a new instance of Binary Tree
   BinaryTree tree = new BinaryTree();
 
   //inserting into the binary tree
   tree.root = new Node(1);
   tree.root.left = new Node(2);
   tree.root.right = new Node(3);
 }
}

```

``` python
4. Binary Tree Program in Python
# Binary Tree in Python
 
# node that hold data, address of the left child, and address of the right child
class Node:def __init__(self, key):# setting left and right child equal to NULL
       self.left = None
       self.right = None
      
       # inserting data into the node
       self.val = key
 
root = Node(1)
root.left = Node(2) 
root.right = Node(3)

```

## Inorder Traversal

Problem Description

Given a binary tree, return the inorder traversal of its nodes' values.



Problem Constraints

1 <= number of nodes <= 105



Input Format

First and only argument is root node of the binary tree, A.



Output Format

Return an integer array denoting the inorder traversal of the given binary tree.



Example Input

Input 1:
```
   1
    \
     2
    /
   3
```
Input 2:
```
   1
  / \
 6   2
    /
   3
```

Example Output

Output 1:

 [1, 3, 2]
Output 2:

 [6, 1, 3, 2]


Example Explanation

Explanation 1:

 The Inorder Traversal of the given tree is [1, 3, 2].
Explanation 2:

 The Inorder Traversal of the given tree is [6, 1, 3, 2].


## Preorder Traversal

Problem Description

Given a binary tree, return the preorder traversal of its nodes values.



Problem Constraints

1 <= number of nodes <= 105



Input Format

First and only argument is root node of the binary tree, A.



Output Format

Return an integer array denoting the preorder traversal of the given binary tree.



Example Input

Input 1:
```
   1
    \
     2
    /
   3
```
Input 2:
```
   1
  / \
 6   2
    /
   3
```

Example Output

Output 1:

 [1, 2, 3]
Output 2:

 [1, 6, 2, 3]



Example Explanation

Explanation 1:

 The Preoder Traversal of the given tree is [1, 2, 3].
Explanation 2:

 The Preoder Traversal of the given tree is [1, 6, 2, 3].


## PostOrder Traversal

Problem Description

Given a binary tree, return the Postorder traversal of its nodes values.



Problem Constraints

1 <= number of nodes <= 105



Input Format

First and only argument is root node of the binary tree, A.



Output Format

Return an integer array denoting the Postorder traversal of the given binary tree.



Example Input

Input 1:
```
   1
    \
     2
    /
   3
```
Input 2:
```
   1
  / \
 6   2
    /
   3
```

Example Output

Output 1:

 [3, 2, 1]
Output 2:

 [6, 3, 2, 1]


Example Explanation

Explanation 1:

 The Preoder Traversal of the given tree is [3, 2, 1].
Explanation 2:

 The Preoder Traversal of the given tree is [6, 3, 2, 1].


## Tree Height

Problem Description

You are given the root node of a binary tree A. You have to find the height of the given tree.

A binary tree's height is the number of nodes along the longest path from the root node down to the farthest leaf node.



Problem Constraints

1 <= Number of nodes in the tree <= 105

0 <= Value of each node <= 109



Input Format

The first and only argument is a tree node A.



Output Format

Return an integer denoting the height of the tree.



Example Input

Input 1:
```
 Values =  1 
          / \     
         4   3   
```

Input 2:
```
 Values =  1      
          / \     
         4   3                       
        /         
       2                                     
```

Example Output

Output 1:

 2 
Output 2:

 3 


Example Explanation

Explanation 1:

 Distance of node having value 1 from root node = 1
 Distance of node having value 4 from root node = 2 (max)
 Distance of node having value 3 from root node = 2 (max)
Explanation 2:

 Distance of node having value 1 from root node = 1
 Distance of node having value 4 from root node = 2
 Distance of node having value 3 from root node = 2
 Distance of node having value 2 from root node = 3 (max)


## Nodes Count

Problem Description

You are given the root node of a binary tree A. You have to find the number of nodes in this tree.



Problem Constraints

1 <= Number of nodes in the tree <= 105

0 <= Value of each node <= 107





Input Format

The first and only argument is a tree node A.



Output Format

Return an integer denoting the number of nodes of the tree.



Example Input

Input 1:
```
 Values =  1 
          / \     
         4   3                        
```
Input 2:

 ```
 Values =  1      
          / \     
         4   3                       
        /         
       2                                     
```

Example Output

Output 1:

 3 
Output 2:

 4 


Example Explanation

Explanation 1:

Clearly, there are 3 nodes 1, 4 and 3.
Explanation 2:

Clearly, there are 4 nodes 1, 4, 3 and 2.

``` go

func solve(A *treeNode )  (int) {
    if A==nil{
        return 0
    }

    return 1+solve(A.left)+solve(A.right)
}

```

## Sum of nodes of a Binary Tree

Problem Description

Given the root of a binary tree A. Return the sum of all the nodes of the binary tree.


Problem Constraints

1 <= Number of nodes in A <= 104

1 <= value of each node <= 104



Input Format

First argument is the root of binary tree A.



Output Format

Return an integer denoting the sum of all the nodes



Example Input

Input 1:
```
 A =   2 
      / \                           
     1   4            
    /                           
   5
```

Input 2:
```
A =    3 
      / \
      6  1
      \   \
       2   7
```

Example Output

Output 1:

12
Output 2:

19


Example Explanation

Explanation 1:

The sum of all the nodes is 12
Explanation 2:

The sum of all the nodes is 19


``` go
func solve(A *treeNode )  (int) {
    if A==nil{
        return 0
    }

    return A.value+solve(A.left)+solve(A.right)
}
```