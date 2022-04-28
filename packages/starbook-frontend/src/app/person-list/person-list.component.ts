import { Component, OnInit } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { HttpService } from '../http.service';
import { Person } from '../models/person';
import { PersonService } from '../person.service';

@Component({
  selector: 'app-person-list',
  templateUrl: './person-list.component.html',
  styleUrls: ['./person-list.component.scss']
})
export class PersonListComponent {

  persons$: Observable<Person[]>;

  constructor(private personService: PersonService) {
    this.persons$ = this.personService.persons$;
  }

}
