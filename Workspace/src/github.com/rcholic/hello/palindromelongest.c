#include <stdio.h>
#include <string.h>

int lps(char *str) {
	int i, j, L, n;
	int **table;

	n = strlen(str);

	//allocate memory for table
	table = (int **) malloc((n+1) * sizeof(int*));

	for (i = 0; i <= n; i++) {
		table[i] = (int *) malloc(n * sizeof(int));
	}

	//initialize the table
	for (i = 0; i <= n; i++) {
		table[i][i] = 1;
	}

	//for every length substrings 
	for (L = 2; L <= n; L++) {
		for (i = 0; i <= n -L + 1; i++) {
			j = i + L - 1;
			if (str[i] == str[j] && L == 2) {
				table[i][j] = 2;
			} else if(str[i] != str[j]) {
				table[i][j] = table[i+1][j] >= table[i][j-1] ? table[i+1][j] : table[i][j-1];

			} else {
				table[i][j] = table[i][j] + 2;
			}

		}
	}
	int maxL = table[0][n-1]

	//clean up memory
	for (i = 0; i <= n; i++) {
		free(table[i]);
	}
	free(table);

	return maxL;
}