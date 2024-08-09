import {
  FormControl,
  FormField,
  FormFieldProps,
  FormLabel,
  FormMessage,
} from "@radix-ui/react-form";
import { Flex, Select, Text, TextField } from "@raystack/apsara";
import React, { CSSProperties } from "react";

import { Control, Controller, UseFormRegister } from "react-hook-form";
import { capitalizeFirstLetter } from "~/utils/helper";
import { MultiSelect } from "./multiselect";

type CustomFieldNameProps = {
  name: string;
  title?: string;
  disabled?: boolean;
  register: UseFormRegister<any>;
  control: Control<any, any>;
  variant?: "textarea" | "input" | "select" | "multiselect";
  style?: CSSProperties;
  options?: Array<{ label: string; value: any }>;
};

export const CustomFieldName = ({
  name,
  title,
  register,
  control,
  disabled = false,
  variant = "input",
  style = {},
  placeholder,
  options = [],
  ...props
}: FormFieldProps &
  CustomFieldNameProps &
  React.RefAttributes<HTMLDivElement>) => {
  const inputTitle = capitalizeFirstLetter(title || name);
  return (
    <FormField
      name={name}
      defaultValue={props.defaultValue}
      style={{ width: "100%" }}
      asChild
    >
      <Flex direction="column" gap="small">
        <Flex
          gap="medium"
          style={{
            alignItems: "baseline",
            justifyContent: "space-between",
          }}
        >
          <FormLabel>
            <Text>{inputTitle}</Text>
          </FormLabel>
          <FormMessage match="valueMissing">
            Please enter your {name}
          </FormMessage>
          <FormMessage match="typeMismatch">
            Please provide a valid {title}
          </FormMessage>
        </Flex>
        <FormControl asChild>
          <Controller
            defaultValue={props.defaultValue}
            name={name}
            control={control}
            render={({ field }) => {
              switch (variant) {
                case "textarea": {
                  return (
                    <textarea
                      {...field}
                      placeholder={
                        placeholder ||
                        `Enter your ${title?.toLowerCase() || name}`
                      }
                      style={style}
                    />
                  );
                }
                case "select": {
                  const { ref, onChange, ...rest } = field;
                  return (
                    <Select
                      {...rest}
                      onValueChange={(value: any) => field.onChange(value)}
                    >
                      <Select.Trigger
                        ref={ref}
                        style={{ height: "26px", width: "100%" }}
                      >
                        <Select.Value placeholder={`Select ${inputTitle}`} />
                      </Select.Trigger>
                      <Select.Content style={{ width: "320px" }}>
                        <Select.Group>
                          {options.map((opt) => (
                            <Select.Item key={opt.value} value={opt.value}>
                              {opt.label}
                            </Select.Item>
                          ))}
                        </Select.Group>
                      </Select.Content>
                    </Select>
                  );
                }
                case "multiselect": {
                  const { ref, onChange, value, ...rest } = field;
                  return (
                    <MultiSelect<string>
                      options={options}
                      onSelect={onChange}
                      selected={value}
                    />
                  );
                }

                default: {
                  return (
                    <TextField
                      {...field}
                      placeholder={
                        placeholder ||
                        `Enter your ${title?.toLowerCase() || name}`
                      }
                      disabled={disabled}
                    />
                  );
                }
              }
            }}
          />
        </FormControl>
      </Flex>
    </FormField>
  );
};

const styles = {
  main: { padding: "32px", width: "80%", margin: 0 },
};
