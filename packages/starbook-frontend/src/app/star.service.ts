import { Injectable } from '@angular/core';
import { filter, merge, Observable, Subject, switchMap } from 'rxjs';
import { HttpService } from './http.service';
import { Star } from './models/star';

@Injectable({
  providedIn: 'root'
})
export class StarService {

  private refreshTrigger = new Subject<number>();

  constructor(private httpClient: HttpService) {

  }

  addStar(comment: string, userId: number) {
    this.httpClient.postStar(comment, userId).subscribe(() => {
      this.refreshTrigger.next(userId);
    })
  }

  starObservable(userId: number): Observable<Star[]> {
    return merge(this.httpClient.fetchPersonStars(userId), this.refreshTrigger.pipe(
      filter(id => id === userId),
      switchMap(() => this.httpClient.fetchPersonStars(userId))))
  }

}
