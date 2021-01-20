import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';

import { AddButtonComponent } from './add-button.component';

@NgModule({
    declarations: [AddButtonComponent],
    imports: [
        CommonModule,
        MatButtonModule,
    ],
    exports: [AddButtonComponent]
})
export class AddButtonModule { }
