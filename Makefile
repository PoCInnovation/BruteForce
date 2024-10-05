
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

install_program:
	echo -n "Do you want to download a default wordlist ?? [y/N] " && read ans && if [ $${ans:-'N'} = 'y' ]; then curl https://raw.githubusercontent.com/drtychai/wordlists/master/dirbuster/directory-list-2.3-medium.txt > default-wordlist.txt; fi

.PHONY: all clean fclean re install_program
