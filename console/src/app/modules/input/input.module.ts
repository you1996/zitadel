import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';

import { InputDirective } from './input.component';

@NgModule({
    imports: [
        CommonModule,
        MatFormFieldModule,
        FormsModule],
    declarations: [InputDirective],
    exports: [InputDirective],
})
export class InputModule { }
