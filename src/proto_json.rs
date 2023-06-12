#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct JsonRecord {
    #[prost(string, tag = "1")]
    pub tag: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub time: ::core::option::Option<::prost_types::Timestamp>,
    #[prost(string, tag = "3")]
    pub json: ::prost::alloc::string::String,
}
