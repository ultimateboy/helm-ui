import { Injectable }    from '@angular/core';
import { Headers, Http } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import { APIURL } from './config/config';


import { Release } from './release';

@Injectable()
export class ReleaseService {

  private reposUrl = APIURL + '/releases';

  constructor(private http: Http) { }

  getReleases(): Promise<Release[]> {
    return this.http.get(this.reposUrl)
               .toPromise()
               .then(response => response.json() as Release[])
               .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getReposSlowly(): Promise<Release[]> {
    return new Promise(resolve => {
      // Simulate server latency with 2 second delay
      setTimeout(() => resolve(this.getReleases()), 2000);
    });
  }

  private headers = new Headers({'Content-Type': 'application/json'});

}
