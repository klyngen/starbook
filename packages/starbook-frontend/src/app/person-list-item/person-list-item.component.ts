import { Component, Input, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { map, Observable } from 'rxjs';
import { Person } from '../models/person';
import { StarService } from '../star.service';

@Component({
  selector: 'app-person-list-item',
  templateUrl: './person-list-item.component.html',
  styleUrls: ['./person-list-item.component.scss']
})
export class PersonListItemComponent implements OnInit {

  @Input()
  person: Person | undefined;

  textField = new FormControl("");

  canSubmit$ = this.textField.valueChanges.pipe(map(value => value?.length > 3));

  starAmount: Observable<number> | undefined;

  constructor(private starService: StarService) {
  }

  ngOnInit(): void {
    if (this.person) {
      this.starAmount = this.starService.starObservable(this.person?.ID).pipe(map(stars => stars.length));
    }
  }

  onSubmitClick() {
    if (this.person) {
      this.starService.addStar(this.textField.value, this.person.ID)
      this.textField.setValue("");
    }
  }
}
