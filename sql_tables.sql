/*
    This script creates the necessary tables for the Configuration Management System.
    It includes a table for storing configuration data with appropriate indexing.

    INSERT INTO Configuration (ClientName, AccessGroups, SoftwarePackages, CustomKey, CustomValue) VALUES (
    'VM_Server1',
    '["admins", "devops"]',
    '["nginx", "dotnet-sdk", "ssm-agent"]',
    NULL,
    NULL
);
*/


IF NOT EXISTS (SELECT * FROM sysobjects WHERE name = 'Configuration' AND xtype = 'U')
CREATE TABLE Configuration (
    ID INT IDENTITY PRIMARY KEY,
    ClientName NVARCHAR(100) NOT NULL,
    AccessGroups NVARCHAR(MAX),       -- comma-separated or JSON
    SoftwarePackages NVARCHAR(MAX),   -- comma-separated or JSON
    CustomKey NVARCHAR(100),          -- optional per-key-value pairs
    CustomValue NVARCHAR(MAX),        -- optional per-key-value pairs
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE()
);

IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'IX_ClientName' AND object_id = OBJECT_ID('Configuration'))
CREATE INDEX IX_ClientName ON Configuration(ClientName);