import { Injectable } from '@angular/core';
import { merge, Observable, startWith, Subject, switchMap, switchMapTo, tap } from 'rxjs';
import { HttpService } from './http.service';
import { Person } from './models/person';

@Injectable({
  providedIn: 'root'
})
export class PersonService {

  persons$: Observable<Person[]>;

  private personTrigger = new Subject();

  constructor(private httpClient: HttpService) {
    //this.persons$ = this.personTrigger.pipe(startWith(this.httpClient.fetchPersons()))
    this.persons$ = merge(
      this.httpClient.fetchPersons(),
      this.personTrigger.pipe(switchMap(() => this.httpClient.fetchPersons()))
    );
  }

  createPerson(name: string, picture: string): Observable<void> {
    return this.httpClient.postPerson(name, picture).pipe(tap(() => this.personTrigger.next(null)));
  }


}
