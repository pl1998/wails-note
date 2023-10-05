export function getDirTree(array) {
    return array.filter((item)=> {
        if(item.is_dir===1) {
            if(item.children===null) {
                return item;
            }else {
                item.children = getDirTree(item.children)
                return  item;
            }
        }
    })
}