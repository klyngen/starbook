import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { map } from 'rxjs';
import { HttpService } from '../http.service';
import { PersonService } from '../person.service';

@Component({
  selector: 'app-create-person',
  templateUrl: './create-person.component.html',
  styleUrls: ['./create-person.component.scss']
})
export class CreatePersonComponent {

  constructor(private personService: PersonService) { }

  nameControl = new FormControl("");
  pictureControl = new FormControl("");

  isValidForm = this.nameControl.valueChanges.pipe(map(value => value?.length > 3));

  onSubmitClick() {
    this.personService.createPerson(this.nameControl.value, this.pictureControl.value).subscribe();
    this.nameControl.setValue("");
    this.pictureControl.setValue("");
  }

}
