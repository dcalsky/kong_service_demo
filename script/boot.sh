cur_dir=$(cd $(dirname $0); pwd)
binary_name=main
conf_dir=$cur_dir/conf/
args="-conf-dir=$conf_dir"

echo "$cur_dir/bin/$binary_name $args"

exec $cur_dir/bin/$binary_name $args