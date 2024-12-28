from tkinter import INSERT


class Node:
    def __init__(self, value):
        self.value = value  # Store the value of the node
        self.left = None    # Pointer to the left child
        self.right = None   # Pointer to the right child

class BinaryTree:
    def __init__(self):
        self.root = None  # Initialize the tree with no root

    def insert(self, value):
        """Insert a value into the binary tree."""
        if self.root is None:
            self.root = Node(value)  # Set the root if tree is empty
        else:
            self._insert_recursive(self.root, value)

    def _insert_recursive(self, current, value):
        """Helper function to insert a value recursively."""
        if value < current.value:
            if current.left is None:
                current.left = Node(value)  # Create a left child node
            else:
                self._insert_recursive(current.left, value)
        else:
            if current.right is None:
                current.right = Node(value)  # Create a right child node
            else:
                self._insert_recursive(current.right, value)

    def inorder_traversal(self):
        """Perform an in-order traversal of the binary tree."""
        result = []
        self._inorder_traversal_recursive(self.root, result)
        return result

    def _inorder_traversal_recursive(self, node, result):
        """Helper function to perform in-order traversal recursively."""
        if node is not None:
            self._inorder_traversal_recursive(node.left, result)
            result.append(node.value)  # Visit the node
            self._inorder_traversal_recursive(node.right, result)

# Example usage:
#tree = BinaryTree()
#tree.insert(10)
#tree.insert(5)
#tree.insert(15)
#tree.insert(3)
#tree.insert(7)

#print("In-order Traversal:", tree.inorder_traversal())


tree = BinaryTree()
array = []
while True:
    inp = input()
    if inp == '': break
    else:
        reportValue = inp.split()
        value = str(reportValue[0]).split(':')
        asli = value[0]
        operations = reportValue[1:]
        for element in operations:
            tree.insert(element)


print("In-order Traversal:", tree.inorder_traversal())