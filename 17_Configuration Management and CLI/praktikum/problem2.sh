#!/bin/sh

my_name=$1
my_fb=$2
my_linkedin=$3

now=$(date)
root="${my_name}\ at\ ${now}"
mkdir "${root}"
cd "${root}"
mkdir -p about_me/personal about_me/professional my_friends my_system_info

fb="https://www.facebook.com/${my_fb}"
linkedin="https://www.linkedin.com/in/${my_linkedin}"
echo $fb > about_me/personal/facebook.txt
echo $linkedin > about_me/professional/linkedin.txt

cd my_friends
wget https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt

cd ..
username=$(whoami)
host=$(uname -a)
echo "My username: ${username}" > my_system_info/about_this_laptop.txt
echo "With host: ${host}" >> my_system_info/about_this_laptop.txt

ping=$(ping -c 3 8.8.8.8)
echo $ping > my_system_info/internet_connection.txt
