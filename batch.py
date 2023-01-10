import csv
import subprocess

filename = "systems.csv"
fields = []

# initializing the titles and rows list
fields = []
rows = []

# reading csv file
with open(filename, 'r') as csvfile:
    # creating a csv reader object
    csvreader = csv.reader(csvfile)

    # extracting field names through first row
    fields = next(csvreader)

    # extracting each data row one by one
    for row in csvreader:
        rows.append(row)

    # get total number of rows
    print("Total no. of rows: %d" % csvreader.line_num)

# printing the field names
print('Field names are: ' + ', '.join(field for field in fields))

# printing first 5 rows
print('\nFirst 5 rows are:\n')
for row in rows[:5]:
    # parsing each column of a row
    for col in row:
        print("%10s" % col, end=" "),
    print('\n')

# run script to create all of them
print('Starting creation script...')

for row in rows[:]:

    name = row[0]
    host = row[1]
    numbers = row[2].replace('#',',')

    do = subprocess.run(["./tbgo", "--host=http://host:port", "--username=", "--password=",
                "--napcreate", "--pbx", "--napproxyhost="+host+":5060", "--numbers="+numbers,
                "--siptransport=voip_sip", "--config=config_1", "--digitmap=digitmap.csv", "--routegroups=55,11",
                "--portrange=Host.pr_voip0", "--customer="+name])

    print(do)

    name = ""
    host = ""
    numbers = ""