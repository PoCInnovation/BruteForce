
NAME	=	bruteforce
SRC		=	src/main.go

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRC)

clean:
	go clean

fclean:
	$(RM) $(NAME)

re: fclean all

.PHONY: all clean fclean re
