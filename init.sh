#!/bin/sh

if [ -e .env ]; then
    echo "Initialize .env has already been completed."
else
    cat > .env << EOF
# Enter your mydns account.
ACC=

# Enter your mydns password.
PASS=
EOF
    echo "Initialize .env has been completed."
fi

echo -e "Edit it as you wish.\n"
