#include <iostream>
#include <stdio.h>
// #include <bits/stdc++.h>
using namespace std;

int main()
{
    int n,test,k,kstart;
    scanf("%d",&test);
    while(test--)
    {
        scanf("%d",&n);
        string ks;
        for(int j=n;j>=0;j--)
        {
            if(j%3==0 && (n-j)%5==0)
            {
                ks="";
                for(int k=0;k<j;k++)
                    ks+='5';
                for(int k=0;k<n-j;k++)
                    ks+='3';
                break;
            }
        }
        if(ks=="")
            cout<<"-1\n";
        else
            cout<<ks<<endl;
    }
    return 0;
}
