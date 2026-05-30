# Forms

Use `leancode_forms`.

## Naming conventions

Use consistent form class names:

- `*FormCubit`
- `*FieldCubit`
- `*Form`
- `*Field`

## Validation behavior

- Define consistent form behavior regarding when the initial validation is executed.
- Define consistent form behavior regarding whether the submit button is enabled if required fields are missing values or some fields have validation errors.

## Form cubits

- Provide form cubits with `BlocProvider`.
- Get form cubits from context during feature cubit creation.
