
NAME	=	bruteforce

SRC		=	src/main.go	\

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRC)

clean:
	go clean

fclean:
	$(RM) $(NAME)

re: fclean all

install_program:
	echo "source $(pwd)/autocompletion/bash/_bruteforce" >> ~/.bashrc
	echo "source $(pwd)/autocompletion/zsh/_bruteforce" >> ~/.zshrc

.PHONY: all clean fclean re install_program
