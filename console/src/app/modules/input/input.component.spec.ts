import { Shallow } from 'shallow-render';

import { InputComponent } from './input.component';
import { InputModule } from './input.module';

describe('InputComponent', () => {
  let shallow: Shallow<InputComponent>;

  beforeEach(() => {
    shallow = new Shallow(InputComponent, InputModule);
  });

  it('should render a given input', async () => {
    const { find } = await shallow.render();
    expect(find('input').nativeElement).toBeTruthy();
  });

  it('should default values', async () => {
    const { instance } = await shallow.render();
    expect(instance.type).toBe('text');
    expect(instance.disabled).toBe(false);
    expect(instance.error).toBe(false);
  });

  it('should render a given input disabled', async () => {
    const { find, instance } = await shallow.render(
      `<bc-input [disabled]=true></bc-input>`,
    );
    expect(instance.disabled).toBe(true);
    expect(
      find('input').nativeElement.attributes['ng-reflect-is-disabled'].value,
    ).toBe('true');
  });

  it('should render a given input readonly', async () => {
    const { find, instance } = await shallow.render(
      `<bc-input [readonly]=true></bc-input>`,
    );
    expect(instance.readonly).toBe(true);
    expect(find('input').nativeElement.attributes['readonly'].value).not.toBe(
      false,
    );
  });

  it('should render a given input error', async () => {
    const { find } = await shallow.render(`<bc-input error="true"></bc-input>`);
    expect(find('input').classes['error']).toBeTruthy();
  });

  it('should render a given input required', async () => {
    const { find } = await shallow.render(
      `<bc-input [required]="true"></bc-input>`,
    );
    expect(
      find('input').nativeElement.attributes['ng-reflect-required'].value,
    ).toBe('true');
  });

  it('should render a given input optional', async () => {
    const { find } = await shallow.render(
      `<bc-input [optional]="true"></bc-input>`,
    );
    expect(find('input').classes['optional']).toBe(true);
  });

  it('should not set optional class when required is set', async () => {
    const { find } = await shallow.render(
      `<bc-input [required]="true" [optional]="true"></bc-input>`,
    );

    expect(find('input').classes['optional']).toBe(false);
  });

  it('should render a given input with maxlength', async () => {
    const { find } = await shallow.render(
      `<bc-input [maxlength]="200"></bc-input>`,
    );
    expect(
      find('input').nativeElement.attributes['ng-reflect-maxlength'].value,
    ).toBe('200');
  });

  it('should render a given input with min, max and step', async () => {
    const { find } = await shallow.render(
      `<bc-input [min]="5" [max]="20" [step]="0.5"></bc-input>`,
    );
    expect(find('input').nativeElement.attributes['min'].value).toBe(
      '5',
    );
    expect(find('input').nativeElement.attributes['max'].value).toBe(
      '20',
    );
    expect(find('input').nativeElement.attributes['step'].value).toBe(
      '0.5',
    );
  });

  it('should render a given placeholder', async () => {
    const { find } = await shallow.render(
      `<bc-input placeholder="foobar"></bc-input>`,
    );
    expect(find('input').nativeElement.placeholder).toBe('foobar');
  });

  it('should emit the value onChange', async () => {
    const { instance, find } = await shallow.render();
    const spy = jest.fn();
    instance.valueChange.subscribe(spy);
    find('input').triggerEventHandler('ngModelChange', null);
    expect(spy).toHaveBeenCalled();
  });

  it('should set focus on load', async () => {
    const { instance, find } = await shallow.render(
      `<bc-input [focusOnLoad]=true></bc-input>`,
    );

    expect(instance.focusOnLoad).toBe(true);
    expect(document.activeElement).toBe(find('input').nativeElement);
  });
});
