## OOPS Java

Problem Description


Construct a class Circle that represents a Circle.

The class should support the following functionalities:-

perimeter() -> returns the perimeter of the circle
area() -> returns the area of the circle

Note: Assume Π (pi) = 3.14 for calculations.

Input Format:

First argument A is an integer representing the number of testcases.
For each case, the radius r is taken as input in new line.
Output Format:

The perimeter and area of the constructed circle is printed.
Sample Input:

2     # number of cases
1     # radius of first case
2     # radius of second case
Sample output:

6.28     #perimeter of first case
3.14     #area of first case
12.56    #perimeter of second case
12.56    #area of second case

``` java

class Circle {
    // Define properties here
    float r;
    
    // Define constructor here
    Circle(float r){
        this.r=r;
    }

    float perimeter(){
		// Complete the function
		return ((float)(2))*(float)3.14*(float)r;
	}
	
	float area(){
		// Complete the function
        return (float)3.14*(r*r);
	}
}

```

## Rectangle

Problem Description

Construct a class Rectangle that represents a rectangle.

The class should support the following functionalities:-

perimeter() -> returns the perimeter of the rectangle
area() -> returns the area of the rectangle
Input format:

First argument A is an integer representing the number of testcases.
For each case, x (length) and y (breadth) are taken as input in new line.
Output format:

The perimeter and area of the constructed rectangle are printed.
Sample Input:

1 # number of cases
1 # radius of first case
2 # radius of second case
Sample output:

4 #perimeter of rectangle
1 #area of rectangle

``` java

class Rectangle {
    // Define properties here
    int length;
    int Breadth;
    
    // Define constructor here
    Rectangle(int length, int Breadth){
        this.length=length;
        this.Breadth=Breadth;
    }

    int perimeter(){
		// Complete the function
		return (2*length + 2*Breadth);
	}
	
	int area(){
		// Complete the function
		return length*Breadth;
	}
}

```

## Class Matrix

Construct a class called Matrix which stores a 2D Array. It should store

The number of rows

The number of columns

The 2D Array itself

Implement the following functionalities inside this class :-

input() -> Reads the input from the user. This method should read the input from the user and populate the entire array. Each row will be in a new line and all the elements in a row will be space-separated.

add(Matrix) -> Returns the sum of two matrices. Assume the matrices provided have the same dimensions.

subtract(Matrix) -> Returns the sum of two matrices. Assume the matrices provided have the same dimensions.

transpose() -> Returns a new matrix containing the transpose of the given original matrix.

print() -> prints the entire matrix row by row. Each row will be in a new line and values in each row should be separated by a single space.

You may define any properties in the class as you see appropriate.

``` java

static class Matrix {
    // Define properties here
int row, column;
	int[][] mat;
	// Define constructor here
	Matrix(int rows, int columns){
        this.row=rows;
        this.column=columns;
        mat=new int[rows][columns];
    }
	
	void input(Scanner sc){
	    // Use the Scanner object passed as argument for taking input and not a new Scanner object
		// Complete the function
		for(int i = 0 ; i < this.row ; i++){
			for(int j = 0 ; j < this.column ; j++){
				mat[i][j] = sc.nextInt();;
			}
		}


	}
	
	Matrix add(Matrix x){
		// Complete the function
		Matrix res = new Matrix(this.row , this.column);
		for(int i = 0 ; i < this.row ; i++){
			for(int j = 0 ; j < this.column ; j++){
				res.mat[i][j] = this.mat[i][j] + x.mat[i][j];
			}
		}
		return res;
	}
	
	Matrix subtract(Matrix x){
		// Complete the function
		Matrix res = new Matrix(this.row , this.column);
		for(int i = 0 ; i < this.row ; i++){
			for(int j = 0 ; j < this.column ; j++){
				res.mat[i][j] = this.mat[i][j] - x.mat[i][j];
			}
		}
		return res;
	}
		
	Matrix transpose(){
		// Complete the function
		Matrix res = new Matrix(this.column , this.row);
		for(int i = 0 ; i < this.row ; i++){
			for(int j = 0 ; j < this.column ; j++){
				res.mat[j][i] = this.mat[i][j];
			}
		}
		return res;
	}
	
	void print(){
		// Complete the function
		for(int i = 0 ; i < this.row ; i++){
			for(int j = 0 ; j < this.column ; j++){
				System.out.print(this.mat[i][j] + " ");
			}
			System.out.print("\n");
		}
	}
}

/*
    Matrix a = new Matrix(10, 5)  // 10 rows, 5 columns
	a.input() 
	Matrix b = new Matrix(10, 5)  // 10 rows, 5 columns
	b.input()
    Matrix c1 = a.add(b)
    Matrix c2 = a.subtract(b)
    Matrix c3 = a.transpose()
    a.print()
*/

```