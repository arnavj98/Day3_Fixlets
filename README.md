# Day3_Fixlets

A basic command-line application written in Go to manage a CSV file. It supports listing, querying, sorting, adding, and deleting entries in a CSV file. The application reads the CSV file, allows users to manipulate its contents through a menu-driven interface, and saves the changes back to the file.

Features

1. List Entries: Displays all entries in the CSV file along with column headers.

2. Query Entries: Filters entries based on a user-specified column and value.

3. Sort Entries: Sorts entries by a chosen column.

4. Add Entry: Adds a new entry by prompting the user for values for each column.

5. Delete Entry: Removes an entry based on the specified row number.

6. Persistent Changes: Saves all updates back to the original CSV file.

Notes

Changes are automatically saved to the CSV file whenever entries are added or deleted.
If the program encounters invalid data, it will skip those rows and continue processing.

Limitations

The application currently assumes the CSV file structure remains constant.
No validation is done to check for duplicate FxiletIDs during addition.
