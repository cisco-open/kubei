import React from 'react';
import classnames from 'classnames';
import { useField } from 'formik';
import { FieldLabel, FieldError } from 'components/Form/utils';

import './text-field.scss';

const TextField = ({className, small=false, label, tooltipText, type="text", ...props}) => {
    const [field, meta] = useField(props);
    
    return (
        <div className={classnames("form-field-wrapper", "text-field", {small}, className)}>
            {!!label && <FieldLabel tooltipId={`form-tooltip-${field.name}`} tooltipText={tooltipText}>{label}</FieldLabel>}
            <input type={type} {...field} className="form-field" />
            {meta.touched && meta.error && <FieldError>{meta.error}</FieldError>}
        </div>
    )
}

export default TextField;