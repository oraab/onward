use std::io;

fn main() {
    println!("Please enter the item type.");
    
    let mut item_type = String::new();
    let mut item_content = String::new();

    io::stdin()
        .read_line(&mut item_type)
        .expect("Failed to read item type.");
    
    println!("Please enter the item.");

    io::stdin()
        .read_line(&mut item_content)
        .expect("Failed to read item.");
   
    let item_content = item_content.trim();
    let item_type = item_type.trim();
 
    println!("You have entered item {item_content} of type {item_type}.");

}  
