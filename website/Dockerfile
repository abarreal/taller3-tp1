FROM mcr.microsoft.com/dotnet/sdk:5.0 as builder
# Everything will be installed under /app.
WORKDIR /app
# Install dependencies.
COPY ./website.csproj .
RUN dotnet restore 
# Copy the rest of the source code into the container and build.
COPY ./ .
RUN dotnet publish -o out --no-restore

# Generate a new image which will only hold the binaries.
FROM mcr.microsoft.com/dotnet/aspnet:5.0
WORKDIR /app
COPY --from=builder /app/out .
CMD ["dotnet", "website.dll"]