
year=$1
day=$2

mkdir -p ./$year/$day
touch ./$year/$day/data_big.dat
touch ./$year/$day/data_small.dat
cp ./template/main.go ./$year/$day

echo
echo "./$year/$day created!"
echo