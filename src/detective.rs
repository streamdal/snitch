use crate::error::CustomError;
use crate::matcher_numeric as numeric;
use crate::matcher_pii as pii;
use crate::{matcher_core as core, FromValue};
use gjson::Value;
use protos::detective::{DetectiveStep, DetectiveType};
use std::str;

pub struct Detective {}

impl Default for Detective {
    fn default() -> Self {
        Detective::new()
    }
}

impl Detective {
    pub fn new() -> Self {
        // env_logger::init();
        Detective {}
    }

    pub fn matches(&self, request: &DetectiveStep) -> Result<bool, CustomError> {
        validate_request(request)?;

        match request
            .type_
            .enum_value()
            .map_err(CustomError::MissingMatchType)?
        {
            DetectiveType::DETECTIVE_TYPE_NUMERIC_EQUAL_TO
            | DetectiveType::DETECTIVE_TYPE_NUMERIC_GREATER_EQUAL
            | DetectiveType::DETECTIVE_TYPE_NUMERIC_GREATER_THAN
            | DetectiveType::DETECTIVE_TYPE_NUMERIC_LESS_EQUAL
            | DetectiveType::DETECTIVE_TYPE_NUMERIC_LESS_THAN => numeric::common(request),

            // Core matchers
            DetectiveType::DETECTIVE_TYPE_STRING_EQUAL => core::string_equal_to(request),
            DetectiveType::DETECTIVE_TYPE_STRING_CONTAINS_ANY => core::string_contains_any(request),
            DetectiveType::DETECTIVE_TYPE_STRING_CONTAINS_ALL => core::string_contains_all(request),
            DetectiveType::DETECTIVE_TYPE_IPV4_ADDRESS
            | DetectiveType::DETECTIVE_TYPE_IPV6_ADDRESS => core::ip_address(request),
            DetectiveType::DETECTIVE_TYPE_REGEX => core::regex(request),
            DetectiveType::DETECTIVE_TYPE_TIMESTAMP_RFC3339 => core::timestamp_rfc3339(request),
            DetectiveType::DETECTIVE_TYPE_TIMESTAMP_UNIX_NANO => core::timestamp_unix_nano(request),
            DetectiveType::DETECTIVE_TYPE_TIMESTAMP_UNIX => core::timestamp_unix(request),
            DetectiveType::DETECTIVE_TYPE_BOOLEAN_FALSE => core::boolean(request, false),
            DetectiveType::DETECTIVE_TYPE_BOOLEAN_TRUE => core::boolean(request, true),
            DetectiveType::DETECTIVE_TYPE_IS_EMPTY => core::is_empty(request),
            DetectiveType::DETECTIVE_TYPE_HAS_FIELD => core::has_field(request),
            DetectiveType::DETECTIVE_TYPE_IS_TYPE => core::is_type(request),
            DetectiveType::DETECTIVE_TYPE_UUID => core::uuid(request),
            DetectiveType::DETECTIVE_TYPE_MAC_ADDRESS => core::mac_address(request),

            // PII matchers
            DetectiveType::DETECTIVE_TYPE_PII_ANY => pii::any(request),
            DetectiveType::DETECTIVE_TYPE_PII_CREDIT_CARD => pii::credit_card(request),
            DetectiveType::DETECTIVE_TYPE_PII_SSN => pii::ssn(request),
            DetectiveType::DETECTIVE_TYPE_PII_EMAIL => pii::email(request),
            DetectiveType::DETECTIVE_TYPE_PII_PHONE => pii::phone(request),
            DetectiveType::DETECTIVE_TYPE_PII_DRIVER_LICENSE => pii::drivers_license(request),
            DetectiveType::DETECTIVE_TYPE_PII_PASSPORT_ID => pii::passport_id(request),
            DetectiveType::DETECTIVE_TYPE_PII_VIN_NUMBER => pii::vin_number(request),
            DetectiveType::DETECTIVE_TYPE_PII_SERIAL_NUMBER => pii::serial_number(request),
            DetectiveType::DETECTIVE_TYPE_PII_LOGIN => pii::login(request),
            DetectiveType::DETECTIVE_TYPE_PII_TAXPAYER_ID => pii::taxpayer_id(request),
            DetectiveType::DETECTIVE_TYPE_PII_ADDRESS => pii::address(request),
            DetectiveType::DETECTIVE_TYPE_PII_SIGNATURE => pii::signature(request),
            DetectiveType::DETECTIVE_TYPE_PII_GEOLOCATION => pii::geolocation(request),
            DetectiveType::DETECTIVE_TYPE_PII_EDUCATION => pii::education(request),
            DetectiveType::DETECTIVE_TYPE_PII_FINANCIAL => pii::financial(request),
            DetectiveType::DETECTIVE_TYPE_PII_HEALTH => pii::health(request),

            DetectiveType::DETECTIVE_TYPE_UNKNOWN => Err(CustomError::Error(
                "match type cannot be unknown".to_string(),
            )),
        }
    }
}

pub fn parse_field<'a, T: FromValue<'a>>(
    data: &'a [u8],
    path: &'a String,
) -> Result<T, CustomError> {
    let data_as_str = str::from_utf8(data)
        .map_err(|e| CustomError::Error(format!("unable to convert bytes to string: {}", e)))?;

    let v = gjson::get(data_as_str, path);

    if !v.exists() {
        return Err(CustomError::Error(format!(
            "path '{}' not found in data",
            path
        )));
    }

    T::from_value(v)
}

fn validate_request(request: &DetectiveStep) -> Result<(), CustomError> {
    match request.type_.enum_value() {
        Ok(value) => {
            if value == DetectiveType::DETECTIVE_TYPE_UNKNOWN {
                return Err(CustomError::MatchError(format!(
                    "unknown match type: {:?}",
                    value
                )));
            }
        }
        Err(value) => {
            return Err(CustomError::MatchError(format!(
                "unexpected match type: {:?}",
                value
            )));
        }
    }

    if request.path.is_empty() {
        return Err(CustomError::Error("path cannot be empty".to_string()));
    }

    if request.input.is_empty() {
        return Err(CustomError::Error("data cannot be empty".to_string()));
    }

    Ok(())
}
