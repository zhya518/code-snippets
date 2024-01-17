public static byte[] toByteArray(final InputStream input) throws IOException {
    final ByteArrayOutputStream output = new ByteArrayOutputStream();
    copy(input, output);
    return output.toByteArray();
}
