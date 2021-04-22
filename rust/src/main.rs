use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input)
        .expect("Couldn't read line");
    
    let inp: u32 = input.trim().parse()
        .expect("Please give me correct string number!");

    pascal_triangle(inp);
}

fn pascal_triangle(rows: u32){
    let mut triangle: Vec<Vec<u32>> = Vec::new();
    for row in 1..(rows+1) {
        let mut curr: Vec<u32> = Vec::new();
        for _ in 0..row {curr.push(0);}
        curr[0] = 1;
        curr[(row-1) as usize] = 1;
        for digit in 1..row-1 {curr[digit as usize] = triangle[(row-2) as usize][(digit-1) as usize] + triangle[(row-2) as usize][digit as usize];}
        println!("{:?}", curr);
        triangle.push(curr);
    }
}
