#include <sys/types.h> 
#include <sys/socket.h> 
#include <stdio.h> 
#include <netinet/in.h>                                                
#include <arpa/inet.h> 
#include <unistd.h> 
#include <stdlib.h>

#define BUFFER_SIZE 40

int main() 
{ 
    char buf[BUFFER_SIZE];
    int client_sockfd; 
    int len; 
    struct sockaddr_in address;// 服务器端网络地址结构体                                           
     int result; 
    client_sockfd = socket(AF_INET, SOCK_STREAM, 0);// 建立客户端socket                               
    address.sin_family = AF_INET; 
    address.sin_addr.s_addr = inet_addr("127.0.0.1");             
    address.sin_port = htons(12000); 
    len = sizeof(address);
    // 与远程服务器建立连接
    result = connect(client_sockfd, (struct sockaddr *)&address, len); 
    if(result<0) 
    { 
         perror("connect failed"); 
         exit(EXIT_FAILURE); 
    } 
    printf("Please input the message:");
    scanf("%s",buf);
    send(client_sockfd,buf,BUFFER_SIZE,0);
    recv(client_sockfd,buf,BUFFER_SIZE,0);
    printf("receive data from server: %s/n",buf);
    close(client_sockfd); 
    return 0; 
}