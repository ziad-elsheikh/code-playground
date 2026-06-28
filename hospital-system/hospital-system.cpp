#include <iostream>
using namespace std;

bool flag = true;

struct patient{
    string name;
    int status;
};

patient db[21][5];

void Add(){
    // user guid
    cout << "Enter specialization, name, status" << endl;

    // take inputs
    int spec, status; string name;
    cin >> spec >> name >> status;

    // add patient
    for (int i = 0; i <= 4; i++){
        if (db[spec][i].name == "") {
            db[spec][i].name = name;
            db[spec][i].status = status;
            cout << "patient added into "<< i << "th slot" << endl; 
            return;
        }
    }
    cout << "we can't add more patients to this specialization" << endl;
    return; 
}

void PrintAllPatients(){
    for (int i = 0; i < 21; i++){
        for (int j = 0; j < 5; j++){
            if (db[i][j].name != "")
            cout << "patient name is: " << db[i][j].name << endl;
        }
    }
}
void NextPatient(){
    // 
    cout << "enter the specialization id" << endl;
    int id; cin >> id;

    // printing
    if (db[id][0].name == "")
        cout << "this specializaition is empty" << endl;
    else
        cout << db[id][0].name << endl;
}

void Exit(){
    flag = false;
}

void Menu(){
    cout << "Enter you choice 1-4:" << endl;
    cout << "   1) add a patient" << endl;
    cout << "   2) print all patients" << endl;
    cout << "   3) get next patient" << endl;
    cout << "   4) Exit" << endl;
}
int getChoice(){
    int input;
    cin >> input;

    // this block added by ai
    // it handels the infinit loop error
    if (cin.fail()) {
        cin.clear();
        cin.ignore(10000, '\n');
        return -1;
    }

    return input;

    
}

void excuteChoice(int i){
    switch (i){
        case 1: Add(); break;
        case 2: PrintAllPatients(); break;
        case 3: NextPatient(); break;
        case 4: Exit(); break;
    }
}
void appLoop(){
    while(flag){
        Menu();
        int choice = getChoice();
        excuteChoice(choice);
    }
}
int main(){
    appLoop();
    return 0;
}