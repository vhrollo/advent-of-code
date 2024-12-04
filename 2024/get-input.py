import aocd
import os
import sys

def main():
    
    if len(sys.argv) != 3:
        print("Usage: get-input.py <day-number> <year>")
        sys.exit(1)

    #args
    day_arg = sys.argv[1]
    year_arg = sys.argv[2]

    try:
        day = int(day_arg)
        year = int(year_arg)
    except ValueError:
        print("Error: Day and year must be integers.")
        sys.exit(1)

    if not 1 <= day <= 25:
        print("Error: Day must be between 1 and 25.")
        sys.exit(1)

    #getting local cookie
    try:
        with open('.cookie.txt', 'r') as f:
            session = f.read().strip()
    except FileNotFoundError:
        print("Error: .cookie.txt not found. Please create this file with your session cookie.")
        sys.exit(1)


    # getting data using aocd
    os.environ["AOC_SESSION"] = session

    try:
        data = aocd.get_data(day=day, year=year)
    except Exception as e:
        print(f"Error fetching data: {e}")
        sys.exit(1)



    #saving in dir
    dir_name = f"day{day:02d}"
    if not os.path.exists(dir_name):
        os.makedirs(dir_name)
        print(f"Created directory: {dir_name}")

    input_file_path = os.path.join(dir_name, 'input.txt')
    try:
        with open(input_file_path, 'w') as f:
            f.write(data)
        print(f"Input data saved to {input_file_path}")
    except IOError as e:
        print(f"Error writing to file: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
